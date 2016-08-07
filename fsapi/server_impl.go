package fsapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path"
	"strings"
	"time"

	"github.com/IronhandedLayman/icfp-origami/objs"
)

type FoldServerBasic struct {
	website     string
	apikey      string
	lastRequest time.Time
	cacheDir    string
	useCache    bool
}

func NewBasicServer(pointWhere string, teamapikey string) FoldServer {
	usr, _ := user.Current()
	return &FoldServerBasic{
		website:     pointWhere,
		apikey:      teamapikey,
		lastRequest: time.Now().Add(-2000 * time.Millisecond),
		cacheDir:    path.Join(usr.HomeDir, ".icfp-origami-cache"),
		useCache:    true,
	}
}

func (fsb *FoldServerBasic) MakeServerRequest(protocol string, cmdNamePath []string, params objs.M, cache bool) (string, error) {
	cacheFilename := path.Join(fsb.cacheDir, strings.Join(cmdNamePath, "-")+".txt")
	//can I grab the cache?
	if fsb.useCache && cache {
		cfs, err := os.Stat(cacheFilename)
		if err == nil && cfs.Size() > 0 && (cmdNamePath[0] != "snapshot" || cfs.ModTime().Hour() == time.Now().Hour()) {
			bans, err := ioutil.ReadFile(cacheFilename)
			return string(bans), err
		}
	}

	//rate limit wait
	waitUntil := fsb.lastRequest.Add(2000 * time.Millisecond)
	if waitUntil.After(time.Now()) {
		time.Sleep(waitUntil.Sub(time.Now()))
	}
	fsb.lastRequest = time.Now()

	cmdName := strings.Join(cmdNamePath, "/")
	client := http.Client{}
	informParams := url.Values{}
	if params != nil {
		for k, v := range params {
			switch v := v.(type) {
			default:
				informParams.Set(k, fmt.Sprintf("%v", v))
			case int:
				informParams.Set(k, fmt.Sprintf("%d", v))
			case string:
				informParams.Set(k, v)
			case time.Time:
				informParams.Set(k, fmt.Sprintf("%d", v.Unix()))
			}
		}
	}
	reqaddr := fmt.Sprintf("http://%s/api/%s", fsb.website, cmdName)
	req, mrerr := http.NewRequest(protocol, reqaddr, strings.NewReader(informParams.Encode()))
	if mrerr != nil {
		return "", fmt.Errorf("CODER ERROR: request malformed")
	}

	req.Header.Set("Expect", "")
	req.Header.Set("X-Api-Key", fsb.apikey)
	if params != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error while requesting: %v", err)
	}

	respBody, rerr := ioutil.ReadAll(resp.Body)
	if rerr != nil {
		return "", fmt.Errorf("Error while reading response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "NOK", fmt.Errorf("Server Error %s while making request: %s", resp.Status, respBody)
	}
	//attempt to write cache file, overwriting what is already there
	ioutil.WriteFile(cacheFilename, respBody, 0666)
	return string(respBody), nil
}

func (fsb *FoldServerBasic) Hello() (string, error) {
	return fsb.MakeServerRequest("GET", []string{"hello"}, nil, false)
}

func (fsb *FoldServerBasic) SnapshotListRequest() (*objs.SnapshotListResponse, error) {
	srvreq, err := fsb.MakeServerRequest("GET", []string{"snapshot", "list"}, nil, true)
	if err != nil {
		return nil, err
	}
	rspObj := objs.SnapshotListResponse{}
	merr := json.Unmarshal(([]byte)(srvreq), &rspObj)
	if merr != nil {
		return nil, merr
	}
	return &rspObj, nil
}

func (fsb *FoldServerBasic) GetBlob(blobHash string) (string, error) {
	return fsb.MakeServerRequest("GET", []string{"blob", blobHash}, nil, true)
}

func (fsb *FoldServerBasic) ProblemSubmission(solutionSpec string, publishTime time.Time) (string, error) {
	return fsb.MakeServerRequest("POST", []string{"problem", "submit"}, objs.M{
		"solution_spec": solutionSpec,
		"publish_time":  publishTime,
	}, false)
}

func (fsb *FoldServerBasic) SolutionSubmission(problemId int, solution string) (string, error) {
	return fsb.MakeServerRequest("POST", []string{"solution", "submit"}, objs.M{
		"problem_id":    problemId,
		"solution_spec": solution,
	}, false)
}

func (fsb *FoldServerBasic) LatestSnapshot() (*objs.Snapshot, error) {
	resp, err := fsb.SnapshotListRequest()
	if err != nil {
		return nil, fmt.Errorf("Error while requesting snapshot list: %v", err)
	}
	if !resp.Ok {
		return nil, fmt.Errorf("Server error while requesting snapshot list: %v", resp)
	}
	var lsh *objs.SnapshotHash
	for _, sh := range resp.Snapshots {
		if lsh == nil || sh.SnapshotTime.After(lsh.SnapshotTime) {
			lsh = sh
		}
	}
	blresp, berr := fsb.GetBlob(lsh.SnapshotHash)
	if berr != nil {
		return nil, fmt.Errorf("Error obtaining snapshot blob: %v", berr)
	}
	snapResp := objs.Snapshot{}
	umerr := json.Unmarshal(([]byte)(blresp), &snapResp)
	if umerr != nil {
		return nil, fmt.Errorf("Error unmarshalling: %v", umerr)
	}
	return &snapResp, nil
}

func (fsb *FoldServerBasic) Scoreboard() ([]objs.UserState, error) {
	snap, err := fsb.LatestSnapshot()
	if err != nil {
		return nil, fmt.Errorf("Could not retrieve snapshot for scoreboard: %v", err)
	}
	uservals := objs.MergeUserData(snap.Users, snap.Leaderboard)
	objs.ByUsers(func(u1, u2 *objs.UserState) bool {
		return u1.Score > u2.Score
	}).Sort(uservals)
	return uservals, nil
}

func (fsb *FoldServerBasic) GetProblemSpec(problemId int) (objs.Problem, error) {
	blresp, err := fsb.LatestSnapshot()
	if err != nil {
		return objs.NoProblem, fmt.Errorf("Error retrieving snapshot: %v", err)
	}
	var nph objs.ProblemHeader
	for _, ph := range blresp.Problems {
		if ph.ProblemId == problemId {
			nph = ph
			break
		}
	}
	if nph.ProblemId == 0 {
		return objs.NoProblem, fmt.Errorf("Could not find problem id %d", problemId)
	}
	pblob, err := fsb.GetBlob(nph.ProblemSpecHash)
	if err != nil {
		return objs.NoProblem, fmt.Errorf("Error retrieving snapshot: %v", err)
	}
	return objs.ParseProblem(pblob)
}
