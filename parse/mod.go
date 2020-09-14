package parse

import (
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/1Password/dep-report/models"
	"github.com/pkg/errors"
	"golang.org/x/mod/modfile"
)

//ParseModules parses the go.mod file and formats the output for further processing
func ParseModules(filepath string) ([]models.Module, error) {
	modBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read go.mod")
	}

	formattedMods, err := modfile.Parse("go.mod", modBytes, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse go.mod")
	}
	var modArray []models.Module

	for _, mod := range formattedMods.Require {
		var tempMod models.Module
		tempMod.Path = mod.Mod.Path
		tempMod.Version = mod.Mod.Version

		modArray = append(modArray, tempMod)
	}
	return modArray, nil
}

//MapModToDependency takes an array of modules as a param and converts it to an []models.dependency
func MapModToDependency(modules []models.Module) []models.Dependency {
	dependencies := make([]models.Dependency, len(modules))
	for i, mod := range modules {
		var dependency models.Dependency
		dependency.Name = cutVersionSuffix(mod.Path)

		// trimming the incompatible flag from the version is necessary to properly
		// find the version tag in github
		mod.Version = strings.TrimSuffix(mod.Version, "+incompatible")

		if strings.Contains(mod.Version, "-") {
			splitVersion := strings.Split(mod.Version, "-")
			dependency.Revision = splitVersion[len(splitVersion)-1]
		} else {
			dependency.Revision = mod.Version
		}
		dependencies[i] = dependency
	}
	return dependencies
}

var majorVersionSuffixRegex = regexp.MustCompile(`/v[0-9]+$`)

// cutting the major version suffix is necessary in order to properly find the repo
// because the repo url does not contain the suffix
func cutVersionSuffix(path string) string {
	if majorVersionSuffixRegex.MatchString(path) {
		splitPath := strings.Split(path, "/")
		return strings.Join(splitPath[:len(splitPath)-1], "/")
	}
	return path
}
