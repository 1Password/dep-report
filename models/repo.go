package models

type license struct {
	Name string `json:"spdx_id"`
}

type LicenseResponse struct {
	License license `json:"license"`
}

type committer struct {
	Date string `json:"date"`
}

type Commit struct {
	Committer committer `json:"committer"`
}

type CommitResponse struct {
	Commit Commit `json:"commit"`
	SHA    string `json:"sha"`
}

type BranchInfo struct {
	Revision string `json:"revision"`
}

type Release struct {
	Name string `json:"tag_name"`
}

type Tag struct {
	Ref string `json:"ref"`
}
