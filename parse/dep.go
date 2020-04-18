package parse

import (
	"dep-report/models"
	"errors"
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

func ReadGopkg(filepath string) (*models.Pkg, error) {
	pkgData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.New("Failed to read" + filepath + err.Error())
	}

	var pkg models.Pkg
	if _, err := toml.Decode(string(pkgData), &pkg); err != nil {
		return nil, errors.New("Failed to json.Unmarshal pkg data" + err.Error())
	}

	return &pkg, nil
}