package report

import (
	"encoding/json"
	"os/exec"
	"strings"
	"time"

	"github.com/1Password/dep-report/models"
	"github.com/1Password/dep-report/versioncontrol"
	"github.com/pkg/errors"
)

const (
	GITHUB  = "github"
	GITLAB  = "gitlab"
	GERRIT  = "gerrit"
	UNKNOWN = "unknown/other"
)

// BuildReport This function is used to create the dependency report
func (g *Generator) BuildReport(productName string, dependencies []models.Dependency) (*models.Report, error) {
	commit, commitTime, err := getCurrentCommitAndCommitTime()
	if err != nil {
		return nil, err
	}

	report := models.Report{
		Product:    productName,
		Commit:     commit,
		CommitTime: commitTime,
		ReportTime: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
	}

	for _, dependency := range dependencies {
		rObj, err := g.reportObjFromDependency(dependency)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create report object from dependency: %v", dependency)
		}

		report.Dependencies = append(report.Dependencies, *rObj)
	}
	return &report, nil
}

// FormatReport takes a report struct and formats it into pretty json
func FormatReport(rawReport models.Report) ([]byte, error) {
	prettyReport, err := json.MarshalIndent(rawReport, "", "  ")
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal indent report")
	}

	return prettyReport, nil
}

func getCurrentCommitAndCommitTime() (string, string, error) {
	commitBytes, err := exec.Command("git", "rev-parse", "HEAD").Output()
	if err != nil {
		return "", "", errors.Wrap(err, "Failed to get current commit")
	}

	commit := strings.TrimSpace(string(commitBytes))

	commitTimeBytes, err := exec.Command("git", "show", "-s", "--format=%cI", "HEAD").Output()
	if err != nil {
		return "", "", errors.Wrap(err, "Failed to get current commit time")
	}

	commitTime := strings.TrimSpace(string(commitTimeBytes))

	return commit, commitTime, nil
}

func (g Generator) reportObjFromDependency(dep models.Dependency) (*models.ReportObject, error) {
	dep.Source = determineSource(dep.Name)

	var reportObject *models.ReportObject
	var err error

	switch dep.Source {
	case GITHUB:
		reportObject, err = versioncontrol.ReportObjFromGithub(dep, g.request)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to generate reportObject from dependency %s", dep.Name)
		}
	case GERRIT:
		reportObject, err = versioncontrol.ReportObjFromGerrit(dep, g.request)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to generate reportObject from dependency %s", dep.Name)
		}
	default:
		reportObject, err = versioncontrol.ReportObjGeneric(dep)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to generate reportObject from dependency %s", dep.Name)
		}

	}

	return reportObject, nil
}

func determineSource(packageName string) string {
	repo := packageName

	if url, ok := versioncontrol.GithubRepoURLForPackage[packageName]; ok {
		repo = url
	}

	switch {
	case strings.Contains(repo, GITHUB):
		if strings.Contains(repo, "repo") {
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
		return UNKNOWN
	}
}
