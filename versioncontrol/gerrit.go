package versioncontrol

import (
	"dep-report/models"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)
//ReportObjFromGerrit uses the data in a dependency object and creates a report object
func ReportObjFromGerrit(dep models.Dependency, r Client) (*models.ReportObject, error){
	var gerritRepoURL string
	var githubRepoURL string

	url, found := GerritRepoURLForPackage[dep.Name]
	if found {
		gerritRepoURL = url
	}
	url, found = GithubRepoURLForPackage[dep.Name]
	if found {
		githubRepoURL = url
	}
	if !found {
		repoName := strings.TrimPrefix(dep.Name, "golang.org/x/")
		gerritRepoURL = "https://go-review.googlesource.com/projects/" + repoName
		githubRepoURL = "https://api.github.com/repos/golang/" + repoName
	}

	reportObject := models.ReportObject{
		Name: dep.Name,
		Website: gerritRepoURL,
		Source: dep.Source,
	}

	//If the dependency comes from go.mod, we have to get the full commit SHA from github before we can call gerrit
	//go.mod returns either semantic version (v0.3.2) or the commit SHA prefix (d3edc9973b7e)
	var githubCommit models.CommitResponse
	if len(dep.Revision) != 40 {
		if err := r.getGithub(githubRepoURL+"/commits/"+dep.Revision, &githubCommit); err != nil {
			return nil, errors.Wrapf(err, "unable to get commit SHA from %s :", githubRepoURL)
		}
		dep.Revision = githubCommit.SHA
	}

	commitURL := gerritRepoURL + "/commits/" + dep.Revision

	var installed models.Commit
	if err := r.getGerrit(commitURL, &installed); err != nil {
		return nil, errors.Wrapf(err, "Unable to get from %s :", commitURL)
	}

	t, err := formatGerritTime(installed.Committer.Date)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to formatGerritTime")
	}
	reportObject.Installed = models.VersionDetails{
		Commit: installed.CommitSHA,
		Time:   t,
	}

	masterURL := gerritRepoURL + "/branches/master"
	var masterInfo models.BranchInfo
	if err := r.getGerrit(masterURL, &masterInfo); err != nil {
		return nil, errors.Wrapf(err, "Unable to get from %s :", masterURL)
	}

	latestURL := gerritRepoURL + "/commits/" + masterInfo.Revision
	var latest models.Commit
	if err := r.getGerrit(latestURL, &latest); err != nil {
		return nil, errors.Wrapf(err, "Unable to get from %s :", latestURL)
	}

	t, err = formatGerritTime(latest.Committer.Date)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to formatGerritTime")
	}
	reportObject.Latest = models.VersionDetails{
		Commit: masterInfo.Revision,
		Time:   t,
	}

	tagsURL := gerritRepoURL + "/tags"
	var tags []models.Tag
	if err := r.getGerrit(tagsURL, &tags); err != nil {
		return nil, errors.Wrapf(err, "Unable to get from %s :", tagsURL)
	}
	if len(tags) > 0 {
		lastTag := tags[len(tags)-1]
		reportObject.Latest.Version = strings.Replace(lastTag.Ref, "refs/tags/", "", 1)
	}

	var ok bool
	reportObject.License, ok = licenseForRepo[dep.Name]
	if !ok {
		reportObject.License = "Unknown license"
	}

	return &reportObject, nil
}

func formatGerritTime(t string) (string, error) {
	t1, err := time.Parse("2006-01-02 15:04:05.999999999", t)
	if err != nil {
		return "", errors.Wrapf(err, "Unable to parse Gerrit time")
	}
	return t1.Format("2006-01-02T15:04:05Z"), nil
}

func (r *Client) getGerrit(url string, target interface{}) error {
	req, err := http.NewRequest("GET", url, nil)

	resp, err := r.HttpClient.Do(req)
	if err != nil {
		return errors.Wrapf(err, "Unable to client.Do")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
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
