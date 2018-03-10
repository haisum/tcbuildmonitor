package build

type change struct {
	Version  string `json:"version"`
	Username string `json:"username"`
	WebURL   string `json:"webUrl"`
}

type buildType struct {
	Name        string `json:"name"`
	ProjectName string `json:"projectName"`
}

type Build struct {
	ID         int    `json:"id"`
	Href       string `json:"href"`
	Status     string `json:"status"`
	WebURL     string `json:"webUrl"`
	StatusText string `json:"statusText"`
	Changes    struct {
		Href string `json:"href"`
	} `json:"changes"`
	Change          []change `json:"change"`
	TestOccurrences struct {
		NewFailed int `json:"newFailed"`
	} `json:"testOccurrences"`
	BuildType buildType `json:"buildType"`
}
