package window

import (
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
)

func RenderEditSerial() g.Widget {
	return g.Custom(func() {
		if ShowEditSerial {
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
		}
	})
}
