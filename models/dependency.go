package models

import "time"

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

type Modules struct {
	Mods []Module
}

type Module struct {
	Path    string    `json:"Path"`
	Version string    `json:"Version"`
	Time    time.Time `json:"Time"`
	Dir     string    `json:"Dir"`
	GoMod   string    `json:"GoMod"`
}
