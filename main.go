package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/1Password/dep-report/models"
	"github.com/1Password/dep-report/parse"
	"github.com/1Password/dep-report/report"
)

const (
	depFilePath   = "/Gopkg.lock"
	goModFilePath = "/go.mod"
)

func main() {
	githubToken := os.Getenv("GITHUB_OAUTH_TOKEN")
	if githubToken == "" {
		log.Fatal("missing argument: GitHub Token")
	}

	productName, ok := os.LookupEnv("DEP_REPORT_PRODUCT")
	_, generateCycloneDX := os.LookupEnv("DEP_REPORT_CYCLONEDX")

	if !ok {
		productName = "b5server"
	}

	dependencies, err := getDependencyFile()
	if err != nil {
		log.Fatalf("unable to parse dependency file: %v", err)
	}

	g := report.NewGenerator(githubToken, productName)

	rawReport, err := g.BuildReport(productName, dependencies)
	if err != nil {
		log.Fatalf("unable to generate report: %v", err)
	}

	var prettyReport []byte

	if generateCycloneDX {
		prettyReport, err = report.FormatCycloneDXReport(*rawReport)
	} else {
		prettyReport, err = report.FormatReport(*rawReport)
	}

	if err != nil {
		log.Fatalf("unable to format report: %v", err)
	}
	fmt.Println(string(prettyReport))
}

func getDependencyFile() ([]models.Dependency, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	var deps []models.Dependency

	switch {
	case fileExists(filepath.Join(wd, depFilePath)):
		pkg, err := parse.ReadGopkg(wd + depFilePath)
		if err != nil {
			return nil, err
		}
		deps = parse.MapPkgToDependency(*pkg)
	case fileExists(filepath.Join(wd, goModFilePath)):
		mods, err := parse.ParseModules(filepath.Join(wd, goModFilePath))
		if err != nil {
			return nil, err
		}
		deps = parse.MapModToDependency(mods)
	}

	return deps, nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
