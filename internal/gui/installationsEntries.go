package gui

import (
	g "github.com/AllenDang/giu"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
)

func InstallationEntries() []g.Widget {
	var widget []g.Widget
	for _, Installations := range software.InstallationsMap {
		for _, installation := range Installations {
			year := g.Label(installation.Year)
			serial := g.Label(installation.License.Serial)
			widget = append(widget, year, serial)
		}
	}
	return widget
}