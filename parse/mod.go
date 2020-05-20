package parse

import (
	"bytes"
	"github.com/1Password/dep-report/models"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)
//ParseModules runs the `go list` command and formats the output for further processing
func ParseModules() ([]models.Module, error) {
	var modArray []models.Module

	cmd := exec.Command("go", "list", "-m", "-mod=mod", "-json", "all")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	goMod, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("unable to execute go list command: %v", err)
	}

	//goMod, err := exec.Command("go", "list", "-m", "-mod=mod", "-json", "all").Output()
	//if err != nil {
	//	return nil, fmt.Errorf("unable to execute go list command: %v", err)
	//}
	goModString := string(goMod)

	splitGoMod := strings.SplitAfter(goModString, "}\n")

	for _, mod := range splitGoMod {
		var module models.Module
		if mod != "" {
			err := json.Unmarshal([]byte(mod), &module)
			if err != nil {
				return nil, fmt.Errorf("unable to unmarshal module into struct: %w", err)
			}
			if module.Version != "" {
				modArray = append(modArray, module)
			}
		}
	}

	return modArray, nil
}
//MapModToDependency takes an array of modules as a param and converts it to an []models.dependency
func MapModToDependency (modules []models.Module) []models.Dependency{
	dependencies := make([]models.Dependency, len(modules))
	for i, mod := range modules {
		var dependency models.Dependency
		dependency.Name = mod.Path

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


