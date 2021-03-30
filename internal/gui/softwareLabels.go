package gui

import (
	g "github.com/AllenDang/giu"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
)

type SoftwareLabels []string

//func getSoftwareWidgetLabels() []g.Widget {
//	var labels []g.Widget
//
//	for key, _ := range software.InstalledSoftwareMap {
//		label := g.Name(key)
//		labels = append(labels, label)
//	}
//	return labels
//}

func GetSoftwareNames() SoftwareLabels {
	var softwareLabels []string
	for key, _ := range software.InstalledSoftwareMap {
		softwareLabels = append(softwareLabels, key)
	}
	return softwareLabels
}

func (softwareLabel SoftwareLabels) ToInterfaceSlice() []interface{} {
	widgetInterface := make([]interface{}, len(softwareLabel))
	for i := range softwareLabel {
		widgetInterface[i] = softwareLabel[i]
	}
	return widgetInterface
}

func GetSoftwareButtons() []g.Widget {
	var buttons []g.Widget

	for key, _ := range software.InstalledSoftwareMap {
		label := g.Label(key)
		buttons = append(buttons, label)
	}
	return buttons
}