package report

import (
	"dep-report/models"
	"dep-report/versioncontrol"
	"flag"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var githubToken = flag.String("githubToken", "", "Token to be used in Github requests")

func TestGenerateReport(t *testing.T) {
	r, c, err := versioncontrol.SetupHTTPRecord("generateReport")
	if err != nil {
		t.Fatalf("unable to setup test recorder: %v", err)
	}
	defer r.Stop()

	g := Generator{
		request: versioncontrol.Client{
			HttpClient: c,
			Token:      *githubToken,
		},
	}

	productName := "dep-report"

	tests := []struct {
		description string
		pkg         []models.Dependency
		wantReport  models.Report
	}{
		{
			description: "Should return report successfully from go.mod",
			pkg: []models.Dependency{
				{
					Name:     "gopkg.in/check.v1",
					Revision: "788fd7840127",
				},
				{
					Name:     "golang.org/x/text",
					Revision: "v0.3.2",
				},
				{
					Name:     "cloud.google.com/go",
					Revision: "v0.54.0",
				},
				{
					Name:     "github.com/xordataexchange/crypt",
					Revision: "b2862e3d0a77",
				},
				{
					Name:     "github.com/pkg/errors",
					Revision: "v0.8.1",
				},
				{
					Name:     "github.com/BurntSushi/toml",
					Revision: "v0.3.1",
				},
			},
			wantReport: models.Report{
				Product:    "dep-report",
				ReportTime: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
				Commit:     "77ae4af8d07bcd816b0f14bdf26cb074f0cfa8b9",
				CommitTime: "2020-04-22T11:02:24-06:00",
				Dependencies: []models.ReportObject{
					{
						Name:    "gopkg.in/check.v1",
						Source:  "github",
						License: "NOASSERTION",
						Website: "https://api.github.com/repos/go-check/check",
						Installed: models.VersionDetails{
							Time:   "2018-06-28T17:31:08Z",
							Commit: "788fd78401277ebd861206a03c884797c6ec5541",
						},
						Latest: models.VersionDetails{
							Time:   "2020-02-27T12:52:54Z",
							Commit: "8fa46927fb4f5b54d48bde78c6c08db205b2298c",
						},
					},
					{
						Name:    "golang.org/x/text",
						Source:  "gerrit",
						License: "BSD-3-Clause",
						Website: "https://go-review.googlesource.com/projects/text",
						Installed: models.VersionDetails{
							Time:   "2019-04-25T21:42:06Z",
							Commit: "342b2e1fbaa52c93f31447ad2c6abc048c63e475",
						},
						Latest: models.VersionDetails{
							Version: "v0.3.2",
							Time:    "2020-03-06T15:41:05Z",
							Commit:  "06d492aade888ab8698aad35476286b7b555c961",
						},
					},
					{
						Name:    "cloud.google.com/go",
						Source:  "gerrit",
						License: "NOASSERTION",
						Website: "https://code-review.googlesource.com/projects/gocloud",
						Installed: models.VersionDetails{
							Time:   "2020-03-05T18:01:17Z",
							Commit: "a6b88cf34a491498e4c7d15c107a31058693e2cb",
						},
						Latest: models.VersionDetails{
							Version: "v0.9.0",
							Time:    "2020-04-23T00:31:42Z",
							Commit:  "c9d3eadce82c530f46cf3c09fc607e329affe4b2",
						},
					},
					{
						Name:    "github.com/xordataexchange/crypt",
						Source:  "github",
						License: "MIT",
						Website: "https://api.github.com/repos/xordataexchange/crypt",
						Installed: models.VersionDetails{
							Time:   "2017-06-26T21:55:01Z",
							Commit: "b2862e3d0a775f18c7cfe02273500ae307b61218",
						},
						Latest: models.VersionDetails{
							Time:   "2017-06-26T21:55:01Z",
							Commit: "b2862e3d0a775f18c7cfe02273500ae307b61218",
						},
					},
					{
						Name:    "github.com/pkg/errors",
						Source:  "github",
						License: "BSD-2-Clause",
						Website: "https://api.github.com/repos/pkg/errors",
						Installed: models.VersionDetails{
							Time:   "2019-01-03T06:52:24Z",
							Commit: "ba968bfe8b2f7e042a574c888954fccecfa385b4",
						},
						Latest: models.VersionDetails{
							Version: "v0.9.1",
							Time:    "2020-01-14T19:47:44Z",
							Commit:  "614d223910a179a466c1767a985424175c39b465",
						},
					},
					{
						Name:    "github.com/BurntSushi/toml",
						Source:  "github",
						License: "MIT",
						Website: "https://api.github.com/repos/BurntSushi/toml",
						Installed: models.VersionDetails{
							Time:   "2018-08-15T10:47:33Z",
							Commit: "3012a1dbe2e4bd1391d42b32f0577cb7bbc7f005",
						},
						Latest: models.VersionDetails{
							Time:   "2018-08-15T10:47:33Z",
							Commit: "3012a1dbe2e4bd1391d42b32f0577cb7bbc7f005",
						},
					},
				},
			},
		},
		{
			description: "Should return report successfully from gopkg",
			pkg: []models.Dependency{
				{
					Name:     "gopkg.in/check.v1",
					Revision: "788fd78401277ebd861206a03c884797c6ec5541",
				},
				{
					Name:     "golang.org/x/text",
					Revision: "342b2e1fbaa52c93f31447ad2c6abc048c63e475",
				},
				{
					Name:     "cloud.google.com/go",
					Revision: "a6b88cf34a491498e4c7d15c107a31058693e2cb",
				},
				{
					Name:     "github.com/xordataexchange/crypt",
					Revision: "b2862e3d0a775f18c7cfe02273500ae307b61218",
				},
				{
					Name:     "github.com/pkg/errors",
					Revision: "645ef00459ed84a119197bfb8d8205042c6df63d",
				},
				{
					Name:     "github.com/BurntSushi/toml",
					Revision: "3012a1dbe2e4bd1391d42b32f0577cb7bbc7f005",
				},
			},
			wantReport: models.Report{
				Product:    "dep-report",
				ReportTime: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
				Commit:     "77ae4af8d07bcd816b0f14bdf26cb074f0cfa8b9",
				CommitTime: "2020-04-22T11:02:24-06:00",
				Dependencies: []models.ReportObject{
					{
						Name:    "gopkg.in/check.v1",
						Source:  "github",
						License: "NOASSERTION",
						Website: "https://api.github.com/repos/go-check/check",
						Installed: models.VersionDetails{
							Time:   "2018-06-28T17:31:08Z",
							Commit: "788fd78401277ebd861206a03c884797c6ec5541",
						},
						Latest: models.VersionDetails{
							Time:   "2020-02-27T12:52:54Z",
							Commit: "8fa46927fb4f5b54d48bde78c6c08db205b2298c",
						},
					},
					{
						Name:    "golang.org/x/text",
						Source:  "gerrit",
						License: "BSD-3-Clause",
						Website: "https://go-review.googlesource.com/projects/text",
						Installed: models.VersionDetails{
							Time:   "2019-04-25T21:42:06Z",
							Commit: "342b2e1fbaa52c93f31447ad2c6abc048c63e475",
						},
						Latest: models.VersionDetails{
							Version: "v0.3.2",
							Time:    "2020-03-06T15:41:05Z",
							Commit:  "06d492aade888ab8698aad35476286b7b555c961",
						},
					},
					{
						Name:    "cloud.google.com/go",
						Source:  "gerrit",
						License: "NOASSERTION",
						Website: "https://code-review.googlesource.com/projects/gocloud",
						Installed: models.VersionDetails{
							Time:   "2020-03-05T18:01:17Z",
							Commit: "a6b88cf34a491498e4c7d15c107a31058693e2cb",
						},
						Latest: models.VersionDetails{
							Version: "v0.9.0",
							Time:    "2020-04-23T00:31:42Z",
							Commit:  "c9d3eadce82c530f46cf3c09fc607e329affe4b2",
						},
					},
					{
						Name:    "github.com/xordataexchange/crypt",
						Source:  "github",
						License: "MIT",
						Website: "https://api.github.com/repos/xordataexchange/crypt",
						Installed: models.VersionDetails{
							Time:   "2017-06-26T21:55:01Z",
							Commit: "b2862e3d0a775f18c7cfe02273500ae307b61218",
						},
						Latest: models.VersionDetails{
							Time:   "2017-06-26T21:55:01Z",
							Commit: "b2862e3d0a775f18c7cfe02273500ae307b61218",
						},
					},
					{
						Name:    "github.com/pkg/errors",
						Source:  "github",
						License: "BSD-2-Clause",
						Website: "https://api.github.com/repos/pkg/errors",
						Installed: models.VersionDetails{
							Time:   "2016-09-29T01:48:01Z",
							Commit: "645ef00459ed84a119197bfb8d8205042c6df63d",
						},
						Latest: models.VersionDetails{
							Version: "v0.9.1",
							Time:    "2020-01-14T19:47:44Z",
							Commit:  "614d223910a179a466c1767a985424175c39b465",
						},
					},
					{
						Name:    "github.com/BurntSushi/toml",
						Source:  "github",
						License: "MIT",
						Website: "https://api.github.com/repos/BurntSushi/toml",
						Installed: models.VersionDetails{
							Time:   "2018-08-15T10:47:33Z",
							Commit: "3012a1dbe2e4bd1391d42b32f0577cb7bbc7f005",
						},
						Latest: models.VersionDetails{
							Time:   "2018-08-15T10:47:33Z",
							Commit: "3012a1dbe2e4bd1391d42b32f0577cb7bbc7f005",
						},
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			gotReport, err := g.GenerateReport(productName, test.pkg)
			if err != nil {
				t.Errorf("GenerateReport failed with errors: %v", err)
			}

			if gotReport != nil {
				gotReport.ReportTime = test.wantReport.ReportTime
				gotReport.CommitTime = test.wantReport.CommitTime
				gotReport.Commit = test.wantReport.Commit
			}
			assert.EqualValues(t, test.wantReport, *gotReport)
		})
	}
}
