package parse

import (
	"testing"
)

func TestReadGopkg(t *testing.T) {
	//filepath := "../Gopkg.lock"
	//wantpkg := &models.Pkg{
	//	Projects: []models.PkgObject{
	//		{
	//			Name:     "github.com/BurntSushi/toml",
	//			Revision: "3012a1dbe2e4bd1391d42b32f0577cb7bbc7f005",
	//		},
	//		{
	//			Name:     "github.com/davecgh/go-spew",
	//			Revision: "8991bc29aa16c548c550c7ff78260e27b9ab7c73",
	//		},
	//		{
	//			Name:     "github.com/dnaeon/go-vcr",
	//			Revision: "b3f5a17c396f1f45e232e36c6eed2577da52d22a",
	//		},
	//		{
	//			Name:     "github.com/pkg/errors",
	//			Revision: "645ef00459ed84a119197bfb8d8205042c6df63d",
	//		},
	//		{
	//			Name:     "github.com/pmezard/go-difflib",
	//			Revision: "792786c7400a136282c1664665ae0a8db921c6c2",
	//		},
	//		{
	//			Name:     "github.com/stretchr/testify",
	//			Revision: "3ebf1ddaeb260c4b1ae502a01c7844fa8c1fa0e9",
	//		},
	//		{
	//			Name:     "golang.org/x/net",
	//			Revision: "d3edc9973b7eb1fb302b0ff2c62357091cea9a30",
	//			Branch:   "master",
	//		},
	//		{
	//			Name:     "golang.org/x/text",
	//			Revision: "342b2e1fbaa52c93f31447ad2c6abc048c63e475",
	//		},
	//		{
	//			Name:     "gopkg.in/yaml.v2",
	//			Revision: "53403b58ad1b561927d19068c655246f2db79d48",
	//		},
	//	},
	//}
	//
	//pkg, err := ReadGopkg(filepath)
	//if err != nil {
	//	t.Errorf("unable to read Gopkg, %v", err)
	//}
	//assert.EqualValues(t, wantpkg, pkg)
}
