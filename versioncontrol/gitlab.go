package versioncontrol

import "dep-report/models"

func ReportObjFromGitlab(r *models.ReportObject, m models.PkgObject) error {
	r.Source = "gitlab"
	// TODO
	return nil
}
