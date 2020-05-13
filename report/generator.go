package report

import (
	"github.com/1Password/dep-report/versioncontrol"
	"net/http"
	"time"
)

//Generator is used to drive the report generation and data retrieval
type Generator struct {
	productName string
	//Client contains details needed to make API calls to github/gerrit/gitlab/etc
	client versioncontrol.Client
}

//NewGenerator creates a Generator struct
func NewGenerator(githubToken string, productName string, slackWebhook string) *Generator {
	generator := Generator{
		productName: productName,
		client: versioncontrol.Client{
			HttpClient: &http.Client{Timeout: 5 * time.Second},
			Token:      githubToken,
			SlackWebhook: slackWebhook,
		},
	}
	return &generator
}
