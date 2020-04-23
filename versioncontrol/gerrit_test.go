package versioncontrol

import (
	"dep-report/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReportObjFromGerrit(t *testing.T) {
	r, c, err := SetupHTTPRecord("reportObjFromGerrit")
	if err != nil {
		t.Fatalf("unable to setup http recorder, %v", err)
	}
	defer r.Stop()

	tests := []struct {
		description      string
		reportObject     *models.ReportObject
		pkgObject        models.PkgObject
		githubToken      string
		wantReportObject *models.ReportObject
	}{
		{
			description: "should successfully return report object when pkgObject is from go.mod and has pseudo version",
			reportObject: &models.ReportObject{
				Name:   "golang.org/x/net",
				Source: "gerrit",
			},
			pkgObject: models.PkgObject{
				Name:     "golang.org/x/net",
				Revision: "d3edc9973b7e",
			},
			githubToken: *githubToken,
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
					Commit: "e086a090c8fdb9982880f0fb6e3db47af1856533",
					Time:   "2020-04-21T23:12:49Z",
				},
			},
		},
		{
			description: "should successfully return report object when pkgObject is from gopkg",
			reportObject: &models.ReportObject{
				Name:   "golang.org/x/net",
				Source: "gerrit",
			},
			pkgObject: models.PkgObject{
				Name:     "golang.org/x/net",
				Revision: "d3edc9973b7eb1fb302b0ff2c62357091cea9a30",
			},
			githubToken: *githubToken,
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
					Commit: "e086a090c8fdb9982880f0fb6e3db47af1856533",
					Time:   "2020-04-21T23:12:49Z",
				},
			},
		},
		{
			description: "should successfully return report object when pkgObject is from go.mod and has semantic version",
			reportObject: &models.ReportObject{
				Name:   "golang.org/x/text",
				Source: "gerrit",
			},
			pkgObject: models.PkgObject{
				Name:     "golang.org/x/text",
				Revision: "v0.3.2",
			},
			githubToken: *githubToken,
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
					Commit: "06d492aade888ab8698aad35476286b7b555c961",
					Time:   "2020-03-06T15:41:05Z",
					Version: "v0.3.2",
				},
			},
		},
		{
			description: "should successfully return report object when pkgObject is from go.mod and has inconsistent url pattern",
			reportObject: &models.ReportObject{
				Name:   "cloud.google.com/go",
				Source: "gerrit",
			},
			pkgObject: models.PkgObject{
				Name:     "cloud.google.com/go",
				Revision: "v0.54.0",
			},
			githubToken: *githubToken,
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
					Commit: "f5bbb0cbd99b043da6d7054a5f495da3ae8dd15b",
					Time:   "2020-04-21T14:37:30Z",
					Version: "v0.9.0",
				},
			},
		},
		{
			description: "should successfully return report object when pkgObject is from go.mod and has semantic version",
			reportObject: &models.ReportObject{
				Name:   "golang.org/x/text",
				Source: "gerrit",
			},
			pkgObject: models.PkgObject{
				Name:     "golang.org/x/text",
				Revision: "v0.3.2",
			},
			githubToken: *githubToken,
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
					Commit: "06d492aade888ab8698aad35476286b7b555c961",
					Time:   "2020-03-06T15:41:05Z",
					Version: "v0.3.2",
				},
			},
		},
		{
			description: "should successfully return report object when pkgObject contains package not in licenseForRepo map",
			reportObject: &models.ReportObject{
				Name:   "golang.org/x/xerrors",
				Source: "gerrit",
			},
			pkgObject: models.PkgObject{
				Name:     "golang.org/x/xerrors",
				Revision: "9bdfabe68543c54f90421aeb9a60ef8061b5b544",
				Branch: "master",
			},
			githubToken: *githubToken,
			wantReportObject: &models.ReportObject{
				Name:    "golang.org/x/xerrors",
				Source:  "gerrit",
				License: "Unknown license",
				Website: "https://go-review.googlesource.com/projects/xerrors",
				Installed: models.VersionDetails{
					Commit: "9bdfabe68543c54f90421aeb9a60ef8061b5b544",
					Time:   "2019-12-04T19:05:36Z",
				},
				Latest: models.VersionDetails{
					Commit: "9bdfabe68543c54f90421aeb9a60ef8061b5b544",
					Time:   "2019-12-04T19:05:36Z",
					Version: "",
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			err := ReportObjFromGerrit(test.reportObject, test.pkgObject, test.githubToken, c)
			if err != nil {
				t.Errorf("unable to compile report for gerrit: %v", err)
			}
			assert.EqualValues(t, test.wantReportObject, test.reportObject)
		})
	}
}

//func TestFormatGerritTime(t *testing.T) {
//	time :=
//}
