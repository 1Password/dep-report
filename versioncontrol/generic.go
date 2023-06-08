package versioncontrol

import "github.com/1Password/dep-report/models"

func ReportObjGeneric(dep models.Dependency) (*models.ReportObject, error) {
	reportObject := models.ReportObject{
		Name:   dep.Name,
		Source: dep.Source,
		Installed: models.VersionDetails{
			Version: dep.Version,
		},
	}
	return &reportObject, nil
}
