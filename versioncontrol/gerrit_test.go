package versioncontrol

import (
	"errors"
	"github.com/1Password/dep-report/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReportObjFromGerrit(t *testing.T) {
	r, c, err := SetupHTTPRecord("reportObjFromGerrit")
	if err != nil {
		t.Fatalf("unable to setup http recorder, %v", err)
	}
	defer r.Stop()

	request := Client{
		HttpClient: c,
		Token: *githubToken,
	}

	tests := []struct {
		description      string
		dependency       models.Dependency
		wantReportObject *models.ReportObject
		wantErr 		 error
	}{
		{
			description: "should successfully return report object when dependency is from go.mod and has pseudo version",
			dependency: models.Dependency{
				Name:     "golang.org/x/net",
				Revision: "d3edc9973b7e",
				Source: "gerrit",
			},
			wantReportObject: &models.ReportObject{
				Name:    "golang.org/x/net",
				Source:  "gerrit",
				License: "BSD-3-Clause",
				Website: "https://go-review.googlesource.com/projects/net",
				Installed: models.VersionDetails{
					Commit: "d3edc9973b7eb1fb302b0ff2c62357091cea9a30",
					Time:   "2020-03-24T14:37:07Z",
				},
				Latest: models.VersionDetails{
					Commit: "ff2c4b7c35a07b0c1e90ce72aa7bfe41bb66a3cb",
					Time:   "2020-04-25T23:01:54Z",
				},
			},
		},
		{
			description: "should successfully return report object when dependency is from gopkg",
			dependency: models.Dependency{
				Name:     "golang.org/x/net",
				Revision: "d3edc9973b7eb1fb302b0ff2c62357091cea9a30",
				Source: "gerrit",
			},
			wantReportObject: &models.ReportObject{
				Name:    "golang.org/x/net",
				Source:  "gerrit",
				License: "BSD-3-Clause",
				Website: "https://go-review.googlesource.com/projects/net",
				Installed: models.VersionDetails{
					Commit: "d3edc9973b7eb1fb302b0ff2c62357091cea9a30",
					Time:   "2020-03-24T14:37:07Z",
				},
				Latest: models.VersionDetails{
					Commit: "ff2c4b7c35a07b0c1e90ce72aa7bfe41bb66a3cb",
					Time:   "2020-04-25T23:01:54Z",
				},
			},
		},
		{
			description: "should successfully return report object when dependency is from go.mod and has semantic version",
			dependency: models.Dependency{
				Name:     "golang.org/x/text",
				Revision: "v0.3.2",
				Source: "gerrit",
			},
			wantReportObject: &models.ReportObject{
				Name:    "golang.org/x/text",
				Source:  "gerrit",
				License: "BSD-3-Clause",
				Website: "https://go-review.googlesource.com/projects/text",
				Installed: models.VersionDetails{
					Commit: "342b2e1fbaa52c93f31447ad2c6abc048c63e475",
					Time:   "2019-04-25T21:42:06Z",
				},
				Latest: models.VersionDetails{
					Commit: "6ca2caf96f159660c33dae334f64e31e5da91752",
					Time:   "2020-04-25T22:59:43Z",
					Version: "v0.3.2",
				},
			},
		},
		{
			description: "should successfully return report object when dependency is from go.mod and has inconsistent url pattern",
			dependency: models.Dependency{
				Name:     "cloud.google.com/go",
				Revision: "v0.54.0",
				Source: "gerrit",
			},
			wantReportObject: &models.ReportObject{
				Name:    "cloud.google.com/go",
				Source:  "gerrit",
				License: "NOASSERTION",
				Website: "https://code-review.googlesource.com/projects/gocloud",
				Installed: models.VersionDetails{
					Commit: "a6b88cf34a491498e4c7d15c107a31058693e2cb",
					Time:   "2020-03-05T18:01:17Z",
				},
				Latest: models.VersionDetails{
					Commit: "35ee6fba71d7166e4e6e68d22182bb590d7e7da2",
					Time:   "2020-04-27T12:16:02Z",
					Version: "v0.9.0",
				},
			},
		},
		{
			description: "should successfully return report object when dependency contains package not in licenseForRepo map",
			dependency: models.Dependency{
				Name:     "golang.org/x/lint",
				Revision: "738671d3881b9731cc63024d5d88cf28db875626",
				Source: "gerrit",
			},
			wantReportObject: &models.ReportObject{
				Name:    "golang.org/x/lint",
				Source:  "gerrit",
				License: "Unknown license",
				Website: "https://go-review.googlesource.com/projects/lint",
				Installed: models.VersionDetails{
					Commit: "738671d3881b9731cc63024d5d88cf28db875626",
					Time:   "2020-03-02T20:58:51Z",
				},
				Latest: models.VersionDetails{
					Commit: "738671d3881b9731cc63024d5d88cf28db875626",
					Time:   "2020-03-02T20:58:51Z",
					Version: "",
				},
			},
			wantErr: errors.New("unable to retrieve license for golang.org/x/lint"),
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			reportObject, err := ReportObjFromGerrit(test.dependency, request)
			if err != nil {
				if test.wantErr != nil {
					assert.EqualError(t, err, test.wantErr.Error())
				} else {
					t.Errorf("unable to compile report for gerrit: %v", err)
				}
			}
			assert.EqualValues(t, test.wantReportObject, reportObject)
		})
	}
}
