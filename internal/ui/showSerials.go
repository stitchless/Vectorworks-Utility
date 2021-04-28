package ui

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
)

// Start by
func init() {
	err := software.GenerateInstalledSoftwareMap()
	if err != nil {
		fmt.Errorf("could not generate installation map for %s: %v", software.AllActiveSoftwareNames, err)
	}
}
// RenderShowSerials shows serials of found supported software
func RenderShowSerials() g.Widget {
	return g.Custom(func() {
		if featureShowSerial == currentFeature {
			for _, softwareName := range software.AllActiveSoftwareNames {
				for _, installation := range software.InstalledSoftwareMap[softwareName] {
					fmt.Println(installation.License.Serial)
				}
			}
		}
	})
}