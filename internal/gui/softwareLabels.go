package gui

import (
	g "github.com/AllenDang/giu"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
)

type SoftwareLabels []string

func getSoftwareWidgetLabels() []g.Widget {
	var labels []g.Widget

	for key, _ := range software.InstallationsMap {
		label := g.Label(key)
		labels = append(labels, label)
	}
	return labels
}

func GetSoftwareNames() SoftwareLabels {
	var softwareLabels []string
	for key, _ := range software.InstallationsMap {
		softwareLabels = append(softwareLabels, key)
	}
	return softwareLabels
}

func (ss SoftwareLabels) ToInterfaceSlice() []interface{} {
	iface := make([]interface{}, len(ss))
	for i := range ss {
		iface[i] = ss[i]
	}
	return iface
}