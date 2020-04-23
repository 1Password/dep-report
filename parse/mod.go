package parse

import (
	"dep-report/models"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

func ParseModules() (*models.Pkg, error) {
	var modArray []models.Module

	goMod, err := exec.Command("go", "list", "-m", "-mod=mod", "-json", "all").Output()
	if err != nil {
		return nil, fmt.Errorf("unable to execute go list command: %v", err)
	}
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

	return mapModToPkg(&models.Modules{
		Mods: modArray,
	}), nil
}

func mapModToPkg (modules *models.Modules) *models.Pkg{
	pkgs := make([]models.PkgObject, len(modules.Mods))
	for i, mod := range modules.Mods {
		var pkg models.PkgObject
		pkg.Name = mod.Path

		if strings.Contains(mod.Version, "-") {
			splitVersion := strings.Split(mod.Version, "-")
			pkg.Revision = splitVersion[len(splitVersion)-1]
		} else {
			pkg.Revision = mod.Version
		}
		pkgs[i] = pkg
	}
	return &models.Pkg{
		Projects: pkgs,
	}
}


