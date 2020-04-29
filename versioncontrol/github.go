package versioncontrol

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

func ReportObjFromGithub(dep models.Dependency, r Client) (*models.ReportObject,error) {
	repoName, err := repoNameFromGithubPackage(dep.Name)
	if err != nil {
		return nil, err
	}

	//consider whether or not we need this var
	repoURL := "https://api.github.com/repos/" + repoName

	reportObject := models.ReportObject{
		Name: dep.Name,
		Website: repoURL,
		Source: dep.Source,
	}

	licenseURL := repoURL + "/license"
	var licenseResponse models.LicenseResponse
	if err := r.getGithub(licenseURL, &licenseResponse); err != nil {
		return nil, errors.Wrapf(err, "Unable to get from %s :", licenseURL)
	}

	reportObject.License = licenseResponse.License.Name

	commitURL := repoURL + "/commits/" + dep.Revision
	var installed models.CommitResponse
	if err := r.getGithub(commitURL, &installed); err != nil {
		return nil, errors.Wrapf(err, "Unable to get from %s :", commitURL)
	}

	reportObject.Installed = models.VersionDetails{
		Commit: installed.SHA,
		Time:   installed.Commit.Committer.Date,
	}

	branchURL := repoURL + "/commits/HEAD"
	var latest models.CommitResponse
	if err := r.getGithub(branchURL, &latest); err != nil {
		return nil, errors.Wrapf(err, "Unable to get from %s :", branchURL)
	}

	reportObject.Latest = models.VersionDetails{
		Commit: latest.SHA,
		Time:   latest.Commit.Committer.Date,
	}

	releaseURL := repoURL + "/releases/latest"
	var release models.Release
	if err := r.getGithub(releaseURL, &release); err != nil {
		return nil, errors.Wrapf(err, "Unable to get from %s :", releaseURL)
	}
	reportObject.Latest.Version = release.Name

	return &reportObject, nil
}

func repoNameFromGithubPackage(packageName string) (string, error) {
	if rawURL, found := GithubRepoURLForPackage[packageName]; found {
		packageName = rawURL
	} else {
		packageName = "https://" + packageName
	}

	u, err := url.Parse(packageName)
	if err != nil {
		return "", fmt.Errorf("unable to parse repo url, %w",err)
	}

	return strings.Replace(u.Path, "/", "", 1), nil
}

func (r *Client) getGithub(url string, target interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return errors.Wrapf(err,"unable to create request for %s", url)
	}
	req.Header.Add("Authorization", "token "+r.Token)

	resp, err := r.HttpClient.Do(req)
	if err != nil {
		return errors.Wrapf(err, "unable to make http request to github")
	}
	if resp.StatusCode == 401 {
		return fmt.Errorf("%s returned from github, verify that GITHUB_OAUTH_TOKEN is set", resp.Status)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "unable to read response body")
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return errors.Wrapf(err, "unable to unmarshal response body to target struct")
	}

	return nil
}
