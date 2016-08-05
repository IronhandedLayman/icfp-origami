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

func (fsb *FoldServerBasic) Hello() (string, error) {
	client := http.Client{}
	reqaddr := fmt.Sprintf("http://%s/api/%s", fsb.website, "hello")
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
