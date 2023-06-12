package report

import (
	"bytes"
	"fmt"

	"github.com/1Password/dep-report/models"
	cdx "github.com/CycloneDX/cyclonedx-go"
	"github.com/gosimple/slug"
)

func FormatCycloneDXReport(rawReport models.Report) ([]byte, error) {
	metadata := cdx.Metadata{
		Component: &cdx.Component{
			BOMRef:  fmt.Sprintf("pkg:golang/1password/%s", slug.Make(rawReport.Product)),
			Type:    cdx.ComponentTypeApplication,
			Name:    rawReport.Product,
			Version: rawReport.Commit,
		},
	}

	components := make([]cdx.Component, len(rawReport.Dependencies))

	for idx, dep := range rawReport.Dependencies {
		components[idx] = cdx.Component{
			BOMRef:  fmt.Sprintf("pkg:golang/%s@%s", dep.Name, dep.Installed.Version),
			Type:    cdx.ComponentTypeLibrary,
			Name:    dep.Name,
			Version: dep.Installed.Version,
			ExternalReferences: &[]cdx.ExternalReference{
				{
					URL:  dep.Website,
					Type: "Repository metadata",
				},
			},
			Licenses: &cdx.Licenses{
				cdx.LicenseChoice{
					License: &cdx.License{
						Name: dep.License,
					},
				},
			},
		}
	}

	bom := cdx.NewBOM()
	bom.Metadata = &metadata
	bom.Components = &components

	var output bytes.Buffer

	err := cdx.NewBOMEncoder(&output, cdx.BOMFileFormatJSON).
		SetPretty(true).Encode(bom)

	if err != nil {
		return nil, err
	}

	return output.Bytes(), nil
}
