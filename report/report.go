package report

import (
	"dep-report/models"
	"dep-report/versioncontrol"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

const (
	GITHUB        = "github"
	GITLAB        = "gitlab"
	GERRIT        = "gerrit"
)

var Client = &http.Client{Timeout: 5 * time.Second}

// GenerateReport This function is used to create the dependency report
func GenerateReport (githubToken string, productName string, pkg *models.Pkg) ([]byte, error){
	commit, commitTime := getCurrentCommitAndCommitTime()

	report := models.Report{
		Product:    productName,
		Commit:     commit,
		CommitTime: commitTime,
		ReportTime: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
	}

	for _, pObj := range pkg.Projects {
		rObj, err := reportObjFromPkgObj(pObj, githubToken)
		if err != nil {
			return nil, fmt.Errorf("failed to create report object from pkg object: %v, %w", pObj, err)
		}

		report.Dependencies = append(report.Dependencies, rObj)
	}

	prettyReport, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("unable to marshal report into pretty json format: %w", err)
	}
	return prettyReport, nil
}

func getCurrentCommitAndCommitTime() (string, string) {
	commitBytes, err := exec.Command("git", "rev-parse", "HEAD").Output()
	if err != nil {
		log.Fatal("Failed to get current commit", err)
	}
	commit := strings.TrimSpace(string(commitBytes))

	commitTimeBytes, err := exec.Command("git", "show", "-s", "--format=%cI", "HEAD").Output()
	if err != nil {
		log.Fatal("Failed to get current commit time", err)
	}
	commitTime := strings.TrimSpace(string(commitTimeBytes))

	return commit, commitTime
}

func reportObjFromPkgObj(m models.PkgObject, githubToken string) (models.ReportObject, error) {
	log.Println("Transforming ", m.Name)
	r := models.ReportObject{
		Name:    m.Name,
		Website: m.Source,
		Source: determineSource(m.Name),
	}

	switch r.Source {
	case GITHUB:
		//TODO fix this when you reorg repo; needs client passed in
		if err := versioncontrol.ReportObjFromGithub(&r, m, githubToken, Client); err != nil {
			return r, err
		}
	case GITLAB:
		if err := versioncontrol.ReportObjFromGitlab(&r, m); err != nil {
			return r, err
		}
	case GERRIT:
		if err := versioncontrol.ReportObjFromGerrit(&r, m, githubToken, Client); err != nil {
			return r, err
		}
	default:
		log.Println("Unable to determine repo source for ", m.Name)
	}

	return r, nil
}

func determineSource(packageName string) string {
	repo := packageName

	if url, ok := versioncontrol.GithubRepoURLForPackage[packageName]; ok {
		repo = url
	}

	if strings.Contains(repo, GITHUB) {
		return GITHUB
	}

	if strings.Contains(repo, "1password.io") {
		return GITLAB
	}

	if strings.Contains(repo, "googlesource") {
		return GERRIT
	}

	if strings.Contains(repo, "golang.org/x") {
		return GERRIT
	}

	return ""
}
