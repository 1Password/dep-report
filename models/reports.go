package models

// Objects used in construction of report
type VersionDetails struct {
	Version string `json:"version,omitempty"`
	Time    string `json:"time"`
	Commit  string `json:"commit"`
}

type ReportObject struct {
	Name      string         `json:"name"`
	Source    string         `json:"source"`
	License   string         `json:"license"`
	Website   string         `json:"website"`
	Installed VersionDetails `json:"installed"`
	Latest    VersionDetails `json:"latest"`
}

type Report struct {
	Product      string         `json:"product"`
	ReportTime   string         `json:"reportTime"`
	Commit       string         `json:"commit"`
	CommitTime   string         `json:"commitTime"`
	Dependencies []ReportObject `json:"dependencies"`
}
