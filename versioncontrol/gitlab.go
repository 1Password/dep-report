package versioncontrol

import "dep-report/models"

func ReportObjFromGitlab(dep models.Dependency) (*models.ReportObject, error) {
	reportObject := models.ReportObject{
		Name: dep.Name,
		Source: dep.Source,
	}
	return &reportObject, nil
}
