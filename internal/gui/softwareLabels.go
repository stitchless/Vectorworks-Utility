package gui

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
)

func SoftwareLabels() []g.Widget {
	var allSoftwareInstallations []software.Installation
	var labels []g.Widget

	for _, s := range software.AllSoftwares {
		installations, err := software.FindInstallationsBySoftware(s)
		if err != nil {
			fmt.Println(err)
		}
		if len(installations) != 0 {
			for _, softwareName := range allSoftwareInstallations {
				label := g.Label(softwareName.Software)
				labels = append(labels, label)
			}
			//allSoftwareInstallations = append(allSoftwareInstallations, installations ...)
		}
	}
	return labels
}
