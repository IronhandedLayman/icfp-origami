package fsapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

func (fsb *FoldServerBasic) MakeServerRequest(cmdName string, params map[string]interface{}) (string, error) {
	client := http.Client{}
	reqaddr := fmt.Sprintf("http://%s/api/%s", fsb.website, cmdName)
	if params != nil {
		fmt.Printf("Set of sent params: %#v", params)
	}
	req, mrerr := http.NewRequest("GET", reqaddr, nil)
	if mrerr != nil {
		return "", fmt.Errorf("CODER ERROR: request malformed")
	}
	req.Header.Set("Expect", "")
	req.Header.Set("X-Api-Key", fsb.apikey)
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
	return fsb.MakeServerRequest("hello", nil)
}

func (fsb *FoldServerBasic) SnapshotListRequest() (string, error) {
	return fsb.MakeServerRequest("snapshot/list", nil)
}
