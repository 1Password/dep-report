package report

import (
	"dep-report/models"
	"dep-report/versionControl"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

const (
	GITHUB        = "github"
	GITLAB        = "gitlab"
	GERRIT        = "gerrit"
)

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
	}

	source := determineSource(m.Name)
	fmt.Println(source)

	switch source {
	case GITHUB:
		if err := versionControl.ReportObjFromGithub(&r, m, githubToken); err != nil {
			return r, err
		}
	case GITLAB:
		if err := versionControl.ReportObjFromGitlab(&r, m); err != nil {
			return r, err
		}
	case GERRIT:
		if err := versionControl.ReportObjFromGerrit(&r, m); err != nil {
			return r, err
		}
	default:
		log.Println("Unable to determine repo source for ", m.Name)
	}

	return r, nil
}

func determineSource(packageName string) string {
	repo := packageName

	if url, ok := versionControl.RepoURLForPackage[packageName]; ok {
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
