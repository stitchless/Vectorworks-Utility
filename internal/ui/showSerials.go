package ui

import (
	g "github.com/AllenDang/giu"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
)

// RenderShowSerials shows serials of found supported software
func RenderShowSerials() g.Widget {
	return g.Custom(func() {
		if featureShowSerial == currentFeature {
			//Render Found installed software buttons.
			for _, softwareName := range software.AllActiveSoftwareNames {
				g.Label(softwareName).Build()
			}
		}
	})
}
