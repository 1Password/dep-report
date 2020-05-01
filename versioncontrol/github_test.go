package versioncontrol

import (
	"github.com/1Password/dep-report/models"
	"flag"
	"github.com/stretchr/testify/assert"
	"testing"
)

var githubToken = flag.String("githubToken", "", "Token to be used in Github requests")

func TestRepoNameFromGithubPackage(t *testing.T) {
	tests := []struct {
		description  string
		packageName  string
		wantRepoName string
	}{
		{
			description:  "Should return repo name when package name is found in map",
			packageName:  "go.opencensus.io",
			wantRepoName: "census-instrumentation/opencensus-go",
		},
		{
			description:  "Should return given packagename param when not found in map",
			packageName:  "github.com/BurntSushi/toml",
			wantRepoName: "BurntSushi/toml",
		},
	}

	for _, test := range tests {
		gotRepoName, err := repoNameFromGithubPackage(test.packageName)
		if err != nil {
			t.Fatalf("unable to get repo name from package: %v", err)
		}
		if gotRepoName != test.wantRepoName {
			t.Errorf("repo name returned did not match expected repo name, want: %s, got: %s", test.wantRepoName, gotRepoName)
		}
	}
}

func TestReportObjFromGithub(t *testing.T) {
	r, c, err := SetupHTTPRecord("reportObjFromGithub")
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
	}{
		{
			description: "should successfully update report object where release is not available",
			dependency: models.Dependency{
				Name:     "github.com/BurntSushi/toml",
				Revision: "3012a1dbe2e4bd1391d42b32f0577cb7bbc7f005",
				Source: "github",
			},
			wantReportObject: &models.ReportObject{
				Name:    "github.com/BurntSushi/toml",
				Source:  "github",
				License: "MIT",
				Website: "https://api.github.com/repos/BurntSushi/toml",
				Installed: models.VersionDetails{
					Commit: "3012a1dbe2e4bd1391d42b32f0577cb7bbc7f005",
					Time:   "2018-08-15T10:47:33Z",
				},
				Latest: models.VersionDetails{
					Commit: "3012a1dbe2e4bd1391d42b32f0577cb7bbc7f005",
					Time:   "2018-08-15T10:47:33Z",
				},
			},
		},
		{
			description: "should successfully update report object where release is available",
			dependency: models.Dependency{
				Name:     "github.com/pkg/profile",
				Revision: "acd64d450fd45fb2afa41f833f3788c8a7797219",
				Source: "github",
			},
			wantReportObject: &models.ReportObject{
				Name:    "github.com/pkg/profile",
				Source:  "github",
				License: "BSD-2-Clause",
				Website: "https://api.github.com/repos/pkg/profile",
				Installed: models.VersionDetails{
					Commit: "acd64d450fd45fb2afa41f833f3788c8a7797219",
					Time:   "2019-11-21T01:09:46Z",
				},
				Latest: models.VersionDetails{
					Commit: "acd64d450fd45fb2afa41f833f3788c8a7797219",
					Time:   "2019-11-21T01:09:46Z",
					Version: "v1.4.0",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			reportObject, err := ReportObjFromGithub(test.dependency, request)
			if err != nil{
				t.Errorf("error returned from ReportObjFromGithub, err: %v", err)
			}
			assert.EqualValues(t, test.wantReportObject, reportObject)
		})
	}
}

func TestGetGithub(t *testing.T) {
	r, c, err := SetupHTTPRecord("getGithub")
	if err != nil {
		t.Fatalf("unable to setup http recorder, %v", err)
	}
	defer r.Stop()

	request := Client{
		HttpClient: c,
		Token: *githubToken,
	}

	var installed models.CommitResponse
	var latest models.CommitResponse
	var license models.LicenseResponse
	tests := []struct {
		description string
		url         string
		target      interface{}
		wantTarget  interface{}
	}{
		{
			description: "getGithub should successfully update latest object",
			url:         "https://api.github.com/repos/BurntSushi/toml/commits/HEAD",
			target:      &latest,
			wantTarget: &models.CommitResponse{
				Commit: models.Commit{
					Committer: models.Committer{
						Date: "2018-08-15T10:47:33Z",
					},
				},
				SHA: "3012a1dbe2e4bd1391d42b32f0577cb7bbc7f005",
			},
		},
		{
			description: "getGithub should successfully get gerrit dependencies",
			url:         "https://api.github.com/repos/golang/net/commits/d3edc9973b7eb1fb302b0ff2c62357091cea9a30",
			target:      &installed,
			wantTarget: &models.CommitResponse{
				Commit: models.Commit{
					Committer: models.Committer{
						Date: "2020-03-24T14:37:07Z",
					},
				},
				SHA: "d3edc9973b7eb1fb302b0ff2c62357091cea9a30",
			},
		},
		{
			//This is not the true installed version of toml, just forcing this for testing purposes.
			description: "getGithub should successfully update installed object",
			url:         "https://api.github.com/repos/BurntSushi/toml/commits/a368813c5e648fee92e5f6c30e3944ff9d5e8895",
			target:      &installed,
			wantTarget: &models.CommitResponse{
				Commit: models.Commit{
					Committer: models.Committer{
						Date: "2017-06-26T11:06:00Z",
					},
				},
				SHA: "a368813c5e648fee92e5f6c30e3944ff9d5e8895",
			},
		},
		{
			description: "getGithub should successfully update license object",
			url:         "https://api.github.com/repos/BurntSushi/toml/license",
			target:      &license,
			wantTarget: &models.LicenseResponse{
				License: models.License{
					Name: "MIT",
				},
			},
		},
		{
			//This test proves that go.mod mapping to gopkg fields works properly
			description: "getGithub should successfully update installed object with semantic version",
			url:         "https://api.github.com/repos/BurntSushi/toml/commits/v0.3.1",
			target:      &latest,
			wantTarget: &models.CommitResponse{
				Commit: models.Commit{
					Committer: models.Committer{
						Date: "2018-08-15T10:47:33Z",
					},
				},
				SHA: "3012a1dbe2e4bd1391d42b32f0577cb7bbc7f005",
			},
		},
		{
			//This test proves that go.mod mapping to gopkg fields works properly
			description: "getGithub should successfully update installed object with commit sha prefix",
			url:         "https://api.github.com/repos/asaskevich/govalidator/commits/475eaeb16496",
			target:      &installed,
			wantTarget: &models.CommitResponse{
				Commit: models.Commit{
					Committer: models.Committer{
						Date: "2020-01-08T20:05:45Z",
					},
				},
				SHA: "475eaeb164960a651e97470412a7d3b0c5036105",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			err := request.getGithub(test.url, &test.target)
			if err != nil {
				t.Errorf("error returned from getGithub: %v", err)
			}
			if test.wantTarget != nil {
				assert.EqualValues(t, test.wantTarget, test.target)
			}
		})
	}
}
