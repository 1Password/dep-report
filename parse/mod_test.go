package parse

import (
	"dep-report/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseModules(t *testing.T) {
	//_, err := ParseModules()
	//if err != nil {
	//	t.Errorf("unable to read go.mod, %v", err)
	//}
}

func TestMapModToPkg(t *testing.T) {
	tests := []struct {
		description string
		modules     []models.Module
		wantPkg     []models.Dependency
	}{
		{
			description: "should handle modules with clean semantic version numbers",
			modules: []models.Module{
				{
					Path:    "github.com/pkg/errors",
					Version: "v0.8.1",
					Time:    time.Now(),
					Dir:     "/home/usr/go/pkg/mod/github.com/pkg/errors@v0.8.1",
					GoMod:   "/home/usr/go/pkg/mod/cache/download/github.com/pkg/errors/@v/v0.8.1.mod",
				},
				{
					Path:    "github.com/BurntSushi/toml",
					Version: "v0.3.1",
					Time:    time.Now(),
					Dir:     "/home/usr/go/pkg/mod/github.com/!burnt!sushi/toml@v0.3.1",
					GoMod:   "/home/usr/go/pkg/mod/cache/download/github.com/!burnt!sushi/toml/@v/v0.3.1.mod",
				},
			},
			wantPkg: []models.Dependency{
				{
					Source:   "",
					Name:     "github.com/pkg/errors",
					Revision: "v0.8.1",
				},
				{
					Source:   "",
					Name:     "github.com/BurntSushi/toml",
					Revision: "v0.3.1",
				},
			},
		},
		{
			description: "should handle modules with pseudo versions",
			modules: []models.Module{
				{
					Path:    "gopkg.in/check.v1",
					Version: "v0.0.0-20161208181325-20d25e280405",
					Time:    time.Now(),
					Dir:     "/home/usr/go/pkg/mod/gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405",
					GoMod:   "/home/usr/go/pkg/mod/cache/download/gopkg.in/check.v1/@v/v0.0.0-20161208181325-20d25e280405.mod",
				},
				{
					Path:    "github.com/xordataexchange/crypt",
					Version: "v0.0.3-0.20170626215501-b2862e3d0a77",
					Time:    time.Now(),
					Dir:     "/home/usr/go/pkg/mod/github.com/xordataexchange/crypt@v0.0.3-0.20170626215501-b2862e3d0a77",
					GoMod:   "/home/usr/go/pkg/mod/cache/download/github.com/xordataexchange/crypt/@v/v0.0.3-0.20170626215501-b2862e3d0a77.mod",
				},
			},
			wantPkg: []models.Dependency{
					{
						Source:   "",
						Name:     "gopkg.in/check.v1",
						Revision: "20d25e280405",
					},
					{
						Source:   "",
						Name:     "github.com/xordataexchange/crypt",
						Revision: "b2862e3d0a77",
					},
				},
			},
		}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got := MapModToDependency(test.modules)
			assert.EqualValues(t, test.wantPkg, got)
		})
	}
}
