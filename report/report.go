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

type Config struct {
	Httpclient *http.Client
	Token string
	Productname string
}

// GenerateReport This function is used to create the dependency report
func (c *Config) GenerateReport (pkg *models.Pkg) ([]byte, error){
	commit, commitTime := getCurrentCommitAndCommitTime()

	report := models.Report{
		Product:    c.Productname,
		Commit:     commit,
		CommitTime: commitTime,
		ReportTime: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
	}

	for _, pObj := range pkg.Projects {
		rObj, err := c.reportObjFromPkgObj(pObj, c.Token)
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

func (c Config) reportObjFromPkgObj(m models.PkgObject, githubToken string) (models.ReportObject, error) {
	log.Println("Transforming ", m.Name)
	r := models.ReportObject{
		Name:    m.Name,
		Website: m.Source,
		Source: determineSource(m.Name),
	}

	switch r.Source {
	case GITHUB:
		if err := versioncontrol.ReportObjFromGithub(&r, m, c.Token, c.Httpclient); err != nil {
			return r, err
		}
	case GITLAB:
		if err := versioncontrol.ReportObjFromGitlab(&r, m); err != nil {
			return r, err
		}
	case GERRIT:
		if err := versioncontrol.ReportObjFromGerrit(&r, m, c.Token, c.Httpclient); err != nil {
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

	switch{
	case strings.Contains(repo, GITHUB):
		if strings.Contains(repo,"repo"){
			return GERRIT
		} else {
			return GITHUB
		}
	case strings.Contains(repo, "1password.io"):
		return GITLAB
	case strings.Contains(repo, "googlesource"):
		return GERRIT
	case strings.Contains(repo, "golang.org/x"):
		return GERRIT
	default:
		return ""
	}
}
