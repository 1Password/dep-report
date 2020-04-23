package main

import (
	"dep-report/models"
	"dep-report/parse"
	"dep-report/report"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
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

	pkg, err := getDependencyFile()
	if err != nil {
		log.Fatalf("unable to parse dependency file: %v",err)
	}

	cfg := report.Config{
		Httpclient: &http.Client{Timeout: 5 * time.Second},
		Token: githubToken,
		Productname: productName,
	}
	prettyReport, err := cfg.GenerateReport(pkg)
	if err != nil {
		log.Fatalf("unable to generate report: %v", err)
	}
	fmt.Println(string(prettyReport))
}

func getDependencyFile() (*models.Pkg, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	var pkg *models.Pkg

	switch {
	case fileExists(wd + depFilePath):
		fmt.Println("Parsing dependencies from gopkg.lock")
		pkg, err = parse.ReadGopkg(wd + depFilePath)
		if err != nil {
			return nil, err
		}
		return pkg, nil
	case fileExists(wd + goModFilePath):
		fmt.Println("Parsing dependencies from go.mod")
		pkg, err = parse.ParseModules()
		if err != nil {
			return nil, err
		}
		return pkg, nil
	}

	return nil, nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
