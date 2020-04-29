package models

import "time"

type Dependency struct {
	//Source is the source of the dependency: github, gitlab, gerrit, etc
	Source string
	//Revision is the commit SHA from gopkg or the version from go.mod.
	//3 formats are used here: full 40 char commit SHA(gopkg), 12 char commit SHA prefix(go.mod) or semantic version number(go.mod)
	Revision string
	//Name is the name of the dependency and comes from Name in gopkg and Path in go.mod
	Name string
}

// PkgObject Objects used when reading from Gopkg.lock
type PkgObject struct {
	Source   string
	Name     string
	VCS      string
	Revision string
	Branch   string
}

//Pkg is a collection of PkgObjects that we use to generate the dependency report
type Pkg struct {
	Projects []PkgObject
}

//Module is a type that represents the json output of `go list`
type Module struct {
	Path    string    `json:"Path"`
	Version string    `json:"Version"`
	Time    time.Time `json:"Time"`
	Dir     string    `json:"Dir"`
	GoMod   string    `json:"GoMod"`
}
