package window

import (
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
	"math"
)

// TODO: create new line
// TODO: Create group
// TODO: Show buttons to navigate the difference features

// RenderTopMenuBar CheckForActiveFeatures gets all active features and returns a slice of software.FeatureName
func RenderTopMenuBar() g.Widget {
	// Button Padding
	const buttonPadding = 20
	numberOfSpaces := len(software.AllActiveFeatures) - 1
	numberOfButtons := len(software.AllActiveFeatures)
	buttonWidth := float32((WindowSize.Width - buttonPadding * numberOfSpaces) / numberOfButtons)
	return g.Custom(func() {
		for i, feature := range software.AllActiveFeatures {
			// Find the x position of the next button
			posX := (buttonWidth * float32(i) + buttonPadding * float32(i)) - (float32(math.Pow(float64(i), float64(numberOfButtons))))
			if i == 0 {
				imgui.SameLineV(0, 0)
			} else {
				imgui.SameLineV(posX, -1)
			}
			g.Button(feature).Size(buttonWidth, 30).OnClick(func() {
				RunModule(feature)
			}).Build()
		}
	})
}