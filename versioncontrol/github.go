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

func ReportObjFromGithub(r *models.ReportObject, m models.PkgObject, githubToken string, c *http.Client) error {
	repoName, err := repoNameFromGithubPackage(m.Name)
	if err != nil {
		return err
	}

	r.Name = m.Name
	//consider whether or not we need this var
	repoURL := "https://api.github.com/repos/" + repoName
	r.Website = repoURL

	licenseURL := repoURL + "/license"
	var licenseResponse models.LicenseResponse
	if err := getGithub(licenseURL, &licenseResponse, githubToken, c); err != nil {
		return errors.Wrapf(err, "Unable to get from %s :", licenseURL)
	}

	r.License = licenseResponse.License.Name

	commitURL := repoURL + "/commits/" + m.Revision
	var installed models.CommitResponse
	if err := getGithub(commitURL, &installed, githubToken, c); err != nil {
		return errors.Wrapf(err, "Unable to get from %s :", commitURL)
	}

	r.Installed = models.VersionDetails{
		Commit: installed.SHA,
		Time:   installed.Commit.Committer.Date,
	}

	branchURL := repoURL + "/commits/HEAD"
	var latest models.CommitResponse
	if err := getGithub(branchURL, &latest, githubToken, c); err != nil {
		return errors.Wrapf(err, "Unable to get from %s :", branchURL)
	}

	r.Latest = models.VersionDetails{
		Commit: latest.SHA,
		Time:   latest.Commit.Committer.Date,
	}

	releaseURL := repoURL + "/releases/latest"
	var release models.Release
	if err := getGithub(releaseURL, &release, githubToken, c); err != nil {
		return errors.Wrapf(err, "Unable to get from %s :", releaseURL)
	}
	r.Latest.Version = release.Name

	return nil
}

func repoNameFromGithubPackage(packageName string) (string, error) {
	if url, found := GithubRepoURLForPackage[packageName]; found {
		packageName = url
	} else {
		packageName = "https://" + packageName
	}

	u, err := url.Parse(packageName)
	if err != nil {
		return "", fmt.Errorf("unable to parse repo url, %w",err)
	}

	return strings.Replace(u.Path, "/", "", 1), nil
}

func getGithub(url string, target interface{}, token string, client *http.Client) error {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "token "+token)

	r, err := client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "unable to make http request to github")
	}
	if r.StatusCode == 401 {
		return fmt.Errorf("%s returned from github, verify that GITHUB_OAUTH_TOKEN is set", r.Status)
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.Wrapf(err, "unable to read response body")
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return errors.Wrapf(err, "unable to unmarshal response body to target struct")
	}

	return nil
}
