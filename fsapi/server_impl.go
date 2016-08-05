package fsapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/IronhandedLayman/icfp-origami/objs"
)

type FoldServerBasic struct {
	website string
	apikey  string
}

func NewBasicServer(pointWhere string, teamapikey string) FoldServer {
	return &FoldServerBasic{
		website: pointWhere,
		apikey:  teamapikey,
	}
}

func (fsb *FoldServerBasic) MakeServerRequest(protocol string, cmdNamePath []string, params objs.M) (string, error) {
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
	return string(respBody), nil
}

func (fsb *FoldServerBasic) Hello() (string, error) {
	return fsb.MakeServerRequest("GET", []string{"hello"}, nil)
}

func (fsb *FoldServerBasic) SnapshotListRequest() (string, error) {
	return fsb.MakeServerRequest("GET", []string{"snapshot", "list"}, nil)
}

func (fsb *FoldServerBasic) GetBlob(blobHash string) (string, error) {
	return fsb.MakeServerRequest("GET", []string{"blob", blobHash}, nil)
}

func (fsb *FoldServerBasic) ProblemSubmission(solutionSpec string, publishTime time.Time) (string, error) {
	return fsb.MakeServerRequest("POST", []string{"problem", "submit"}, objs.M{
		"solution_spec": solutionSpec,
		"publish_time":  publishTime,
	})
}

func (fsb *FoldServerBasic) SolutionSubmission(problemId int, solution string) (string, error) {
	return fsb.MakeServerRequest("POST", []string{"solution", "submit"}, objs.M{
		"problem_id":    problemId,
		"solution_spec": solution,
	})
}