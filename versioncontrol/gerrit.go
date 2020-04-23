package versioncontrol

import (
	"dep-report/models"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func ReportObjFromGerrit(r *models.ReportObject, m models.PkgObject, token string, c *http.Client) error {
	r.Name = m.Name

	var gerritRepoURL string
	var githubRepoURL string

	url, found := GerritRepoURLForPackage[r.Name]
	if found {
		gerritRepoURL = url
	}
	url, found = GithubRepoURLForPackage[r.Name]
	if found {
		githubRepoURL = url
	}
	if !found {
		repoName := strings.TrimPrefix(r.Name, "golang.org/x/")
		gerritRepoURL = "https://go-review.googlesource.com/projects/" + repoName
		githubRepoURL = "https://api.github.com/repos/golang/" + repoName
	}

	r.Website = gerritRepoURL

	//If the pkgObject comes from go.mod, we have to get the full commit SHA from github before we can call gerrit
	//go.mod returns either semantic version (v0.3.2) or the commit SHA prefix (d3edc9973b7e)
	var githubCommit models.CommitResponse
	if len(m.Revision) != 40 {
		if err := getGithub(githubRepoURL+"/commits/"+m.Revision, &githubCommit, token, c); err != nil {
			return errors.Wrapf(err, "unable to get commit SHA from %s :", githubRepoURL)
		}
		m.Revision = githubCommit.SHA
	}

	commitURL := gerritRepoURL + "/commits/" + m.Revision

	var installed models.Commit
	if err := getGerrit(commitURL, &installed, c); err != nil {
		return errors.Wrapf(err, "Unable to get from %s :", commitURL)
	}

	t, err := formatGerritTime(installed.Committer.Date)
	if err != nil {
		return errors.Wrapf(err, "Unable to formatGerritTime")
	}
	r.Installed = models.VersionDetails{
		Commit: installed.CommitSHA,
		Time:   t,
	}

	masterURL := gerritRepoURL + "/branches/master"
	var masterInfo models.BranchInfo
	if err := getGerrit(masterURL, &masterInfo, c); err != nil {
		return errors.Wrapf(err, "Unable to get from %s :", masterURL)
	}

	latestURL := gerritRepoURL + "/commits/" + masterInfo.Revision
	var latest models.Commit
	if err := getGerrit(latestURL, &latest, c); err != nil {
		return errors.Wrapf(err, "Unable to get from %s :", latestURL)
	}

	t, err = formatGerritTime(latest.Committer.Date)
	if err != nil {
		return errors.Wrapf(err, "Unable to formatGerritTime")
	}
	r.Latest = models.VersionDetails{
		Commit: masterInfo.Revision,
		Time:   t,
	}

	tagsURL := gerritRepoURL + "/tags"
	var tags []models.Tag
	if err := getGerrit(tagsURL, &tags, c); err != nil {
		return errors.Wrapf(err, "Unable to get from %s :", tagsURL)
	}
	if len(tags) > 0 {
		lastTag := tags[len(tags)-1]
		r.Latest.Version = strings.Replace(lastTag.Ref, "refs/tags/", "", 1)
	}

	var ok bool
	r.License, ok = licenseForRepo[m.Name]
	if !ok {
		r.License = "Unknown license"
		fmt.Printf("License info for %s not provided", m.Name)
	}

	return nil
}

func formatGerritTime(t string) (string, error) {
	t1, err := time.Parse("2006-01-02 15:04:05.999999999", t)
	if err != nil {
		return "", errors.Wrapf(err, "Unable to parse Gerrit time")
	}
	return t1.Format("2006-01-02T15:04:05Z"), nil
}

func getGerrit(url string, target interface{}, c *http.Client) error {
	req, err := http.NewRequest("GET", url, nil)

	r, err := c.Do(req)
	if err != nil {
		return errors.Wrapf(err, "Unable to client.Do")
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.Wrapf(err, "Unable to ioutil.ReadAll")
	}

	// Gerrit REST API appends a magic string before json body which needs to be removed
	// https://gerrit-review.googlesource.com/Documentation/rest-api.html#output
	bodyString := strings.Replace(string(body), ")]}'\n", "", 1)

	err = json.Unmarshal([]byte(bodyString), target)
	if err != nil {
		return errors.Wrapf(err, "Unable to json.Unmarshal")
	}

	return nil
}
