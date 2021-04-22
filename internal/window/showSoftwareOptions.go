package window

import (
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
)

// RenderActiveSoftwareButtons will generate the UI needed to display all found software installations
func RenderActiveSoftwareButtons() g.Widget {
	return g.Custom(func() {
		// Render Demo Button
		g.Button("Demo").Size(80, BtnHeight).OnClick(func() {
			ShowDemoWindow = true
		}).Build()

		// Render Found installed software buttons.
		for i, softwareName := range software.AllActiveSoftwareNames {
			if _, ok := software.InstalledSoftwareMap[softwareName]; ok {
				if i == 0 {
					// Set first button location
					imgui.SameLineV(softwareBtnPosX(), 0)
				} else {
					// Offset from previous location
					imgui.SameLineV(softwareBtnPosX()+BtnWidth+BtnPadding, 0)
				}

				// Show buttons from all found installed software
				g.Button(softwareName).Size(BtnWidth, BtnHeight).OnClick(func() {
					ActiveSoftwareTab = softwareName
				}).Build()
			}
		}

		// Show the exit button (X) on the right side.
		imgui.SameLineV(float32(WindowSize.Width)-BtnHeight-float32(10), -1)
		g.Button("X").Size(BtnHeight, BtnHeight).OnClick(onQuit).Build()
	})
}