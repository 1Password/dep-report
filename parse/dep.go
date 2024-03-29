package parse

import (
	"io/ioutil"

	"github.com/1Password/dep-report/models"
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

// ReadGopkg takes a filepath and reads the gopkg file at that filepath.
func ReadGopkg(filepath string) (*models.Pkg, error) {
	pkgData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to read file at filepath: %s", filepath)
	}

	var pkg models.Pkg
	if err := toml.Unmarshal(pkgData, &pkg); err != nil {
		return nil, errors.Wrap(err, "Failed to json.Unmarshal pkg data")
	}

	return &pkg, nil
}

// MapPkgToDependency maps gopkg fields to the models.Dependency struct
func MapPkgToDependency(packages models.Pkg) []models.Dependency {
	dependencies := make([]models.Dependency, len(packages.Projects))
	for i, pkg := range packages.Projects {
		dependency := models.Dependency{
			Source:   "Gopkg",
			Revision: pkg.Revision,
			Name:     pkg.Name,
			Version:  "unavailable: installed as Gopkg",
		}
		dependencies[i] = dependency
	}
	return dependencies
}
