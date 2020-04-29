package report

import (
	"dep-report/versioncontrol"
	"net/http"
	"time"
)

//Generator is used to drive the report generation and data retrieval
type Generator struct {
	//Client contains details needed to make API calls to github/gerrit/gitlab/etc
	request versioncontrol.Client
}

//NewGenerator creates a Generator struct
func NewGenerator(githubToken string, productName string) *Generator {
	generator := Generator{
		request: versioncontrol.Client{
			HttpClient: &http.Client{Timeout: 5 * time.Second},
			Token:      githubToken,
		},
	}
	return &generator
}
