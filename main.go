package main

import (
	"fmt"
	"github.com/1Password/dep-report/models"
	"github.com/1Password/dep-report/parse"
	"github.com/1Password/dep-report/report"
	"log"
	"os"
	"path/filepath"
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

	if !ok {
		productName = "b5server"
	}

	//This env variable is required in order for us to report license failures to slack
	slackWebhookURL := os.Getenv("DEP_WEBHOOK")
	if slackWebhookURL == "" {
		log.Fatal("missing env variable: DEP_WEBHOOK")
	}

	dependencies, err := getDependencyFile()
	if err != nil {
		log.Fatalf("unable to parse dependency file: %v", err)
	}

	g := report.NewGenerator(githubToken, productName, slackWebhookURL)

	rawReport, err := g.BuildReport(dependencies)
	if err != nil {
		log.Fatalf("unable to generate report: %v", err)
	}

	prettyReport, err := report.FormatReport(*rawReport)
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
		mods, err := parse.ParseModules()
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
