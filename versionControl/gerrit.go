package versionControl

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

func ReportObjFromGerrit(r *models.ReportObject, m models.PkgObject) error {
	r.Source = "gerrit"
	r.Name = m.Name

	var repoURL string
	if url, found := RepoURLForPackage[r.Name]; found {
		repoURL = url
	} else {
		repoName := strings.TrimPrefix(r.Name, "golang.org/x/")
		repoURL = "https://go-review.googlesource.com/projects/" + repoName
	}
	r.Website = repoURL

	commitURL := repoURL + "/commits/" + m.Revision
	var installed models.Commit
	if err := getGerrit(commitURL, &installed); err != nil {
		return errors.Wrapf(err, "Unable to get from %s :", commitURL)
	}

	t, err := formatGerritTime(installed.Committer.Date)
	if err != nil {
		return errors.Wrapf(err, "Unable to formatGerritTime")
	}
	r.Installed = models.VersionDetails{
		Commit: m.Revision,
		Time:   t,
	}

	masterURL := repoURL + "/branches/master"
	var masterInfo models.BranchInfo
	if err := getGerrit(masterURL, &masterInfo); err != nil {
		return errors.Wrapf(err, "Unable to get from %s :", masterURL)
	}

	latestURL := repoURL + "/commits/" + masterInfo.Revision
	var latest models.Commit
	if err := getGerrit(latestURL, &latest); err != nil {
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

	tagsURL := repoURL + "/tags"
	tags := []models.Tag{}
	if err := getGerrit(tagsURL, &tags); err != nil {
		return errors.Wrapf(err, "Unable to get from %s :", tagsURL)
	}
	if len(tags) > 0 {
		lastTag := tags[len(tags)-1]
		r.Latest.Version = strings.Replace(lastTag.Ref, "refs/tags/", "", 1)
	}

	var ok bool
	r.License, ok = licenseForRepo[m.Name]
	if !ok {
		return fmt.Errorf("License info for %s not provided", m.Name)
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

func getGerrit(url string, target interface{}) error {
	req, err := http.NewRequest("GET", url, nil)

	r, err := Client.Do(req)
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
