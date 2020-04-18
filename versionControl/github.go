package versionControl

import (
	"dep-report/models"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func ReportObjFromGithub(r *models.ReportObject, m models.PkgObject, githubToken string) error {
	r.Source = "github"
	repoName, err := repoNameFromGithubPackage(m.Name)
	if err != nil {
		return err
	}
	fmt.Println(repoName)
	r.Name = m.Name
	//consider whether or not we need this var
	repoURL := "https://api.github.com/repos/" + r.Name
	r.Website = repoURL

	licenseURL := repoURL + "/license"
	var licenseResponse models.LicenseResponse
	if err := GetGithub(licenseURL, &licenseResponse, githubToken); err != nil {
		return errors.Wrapf(err, "Unable to get from %s :", licenseURL)
	}

	r.License = licenseResponse.License.Name

	commitURL := repoURL + "/commits/" + m.Revision
	var installed models.CommitResponse
	if err := GetGithub(commitURL, &installed, githubToken); err != nil {
		return errors.Wrapf(err, "Unable to get from %s :", commitURL)
	}

	r.Installed = models.VersionDetails{
		Commit: m.Revision,
		Time:   installed.Commit.Committer.Date,
	}

	branchURL := repoURL + "/commits/HEAD"
	var latest models.CommitResponse
	if err := GetGithub(branchURL, &latest, githubToken); err != nil {
		return errors.Wrapf(err, "Unable to get from %s :", branchURL)
	}

	r.Latest = models.VersionDetails{
		Commit: latest.SHA,
		Time:   latest.Commit.Committer.Date,
	}

	releaseURL := repoURL + "/releases/latest"
	var release models.Release
	if err := GetGithub(releaseURL, &release, githubToken); err != nil {
		return errors.Wrapf(err, "Unable to get from %s :", releaseURL)
	}
	r.Latest.Version = release.Name

	return nil
}

func repoNameFromGithubPackage(packageName string) (string, error) {
	if url, found := RepoURLForPackage[packageName]; found {
		packageName = url
	}

	u, err := url.Parse(packageName)
	if err != nil {
		return "", fmt.Errorf("unable to parse repo url, %w",err)
	}

	return strings.Replace(u.Path, "/", "", 1), nil
}

func GetGithub(url string, target interface{}, token string) error {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "token "+token)

	r, err := Client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "Unable to client.Do")
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.Wrapf(err, "Unable to ioutil.ReadAll")
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return errors.Wrapf(err, "Unable to json.Unmarshal")
	}

	return nil
}
