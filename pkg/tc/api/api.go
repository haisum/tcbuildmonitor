package api

import (
	"encoding/json"
	"fmt"
	"github.com/haisum/tcbuildmonitor/pkg/tc/build"
	"io/ioutil"
	"net/http"
)

const (
	BuildsEndpoint = "/app/rest/builds?locator=buildType:(id:%s)&state=finished"
	SessionName    = "TCSESSIONID"
)

type doer interface {
	Do(req *http.Request) (*http.Response, error)
}

type tcAPI struct {
	Client    doer
	TcURL     string
	Username  string
	Password  string
	SessionID string
}

func New(client doer, url, username, password string) *tcAPI {
	return &tcAPI{
		client,
		url,
		username,
		password,
		"",
	}
}

func (tc *tcAPI) Get(buildTypeID string) ([]build.Build, error) {
	builds, err := tc.loadBuilds(buildTypeID)
	if err != nil {
		return builds, err
	}
	for index := range builds {
		err = tc.loadBuild(&builds[index])
		if err != nil {
			return builds, err
		}
		err = tc.loadChanges(&builds[index])
		if err != nil {
			return builds, err
		}
	}
	return builds, nil
}

func (tc *tcAPI) loadBuilds(buildTypeID string) ([]build.Build, error) {
	var builds struct {
		Builds []build.Build `json:"build"`
	}
	err := tc.load(fmt.Sprintf("%s"+BuildsEndpoint, tc.TcURL, buildTypeID), &builds)
	return builds.Builds, err
}

func (tc *tcAPI) loadBuild(b *build.Build) error {
	return tc.load(fmt.Sprintf("%s"+b.Href, tc.TcURL), b)
}

func (tc *tcAPI) loadChanges(b *build.Build) error {
	return tc.load(fmt.Sprintf("%s"+b.Changes.Href, tc.TcURL), b)
}

func (tc *tcAPI) load(url string, v interface{}) error {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	request.Header.Set("Accept", "application/json")
	if tc.SessionID != "" {
		request.AddCookie(&http.Cookie{
			Name:  SessionName,
			Value: tc.SessionID,
		})
	} else {
		request.SetBasicAuth(tc.Username, tc.Password)
	}
	response, err := tc.Client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	// unauthorized due to cookie?
	if response.StatusCode == 401 && tc.SessionID != "" {
		tc.SessionID = ""
		return tc.load(url, v)
	}
	if response.StatusCode >= 400 {
		body, _ := ioutil.ReadAll(response.Body)
		return fmt.Errorf("http error code: %d, body: %s", response.StatusCode, body)
	}
	for _, cookie := range response.Cookies() {
		if cookie.Name == SessionName {
			tc.SessionID = cookie.Value
		}
	}
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(v)
	return err
}
