package models

type License struct {
	Name string `json:"spdx_id"`
}

type LicenseResponse struct {
	License License `json:"License"`
}

type Committer struct {
	Date string `json:"date"`
}

type Commit struct {
	Committer Committer `json:"Committer"`
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
