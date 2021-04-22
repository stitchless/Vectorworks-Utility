package window

import (
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
)

// TODO: create new line
// TODO: Create group
// TODO: Show buttons to navigate the difference features

// CheckForActiveFeatures gets all active features and returns a slice of software.FeatureName
func RenderTopMenuBar() g.Widget {
	buttonWidth := float32(WindowSize.Width / len(software.AllActiveFeatures))
	return g.Custom(func() {
		for i, feature := range software.AllActiveFeatures {
			if i == 0 {
				// Set first button location
				imgui.SameLineV(-1, 0)
			} else {
				// Offset from previous location
				imgui.SameLineV(buttonWidth, 0)
			}
			g.Button(feature).Size(buttonWidth, 30).OnClick(func() {
				RunModule(feature)
			}).Build()
		}
	})
}