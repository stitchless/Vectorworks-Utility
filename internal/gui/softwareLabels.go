package gui

import (
	g "github.com/AllenDang/giu"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
)

func SoftwareLabels() []g.Widget {
	var labels []g.Widget

	for key, _ := range software.InstallationsMap {
		label := g.Label(key)
		labels = append(labels, label)
	}
	return labels
}

func GetSoftwareNames() []string {
	var softwareLabels []string
	for key, _ := range software.InstallationsMap {
		softwareLabels = append(softwareLabels, key)
	}
	return softwareLabels
}