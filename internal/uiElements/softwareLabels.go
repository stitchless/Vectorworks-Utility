package uiElements

import (
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
	"os"
)

const (
	WindowW int = 800
	WindowH int = 600
)

// softwareBtnPosX will calculate the middle of the window minus half the UI elements drawn to the line
func softwareBtnPosX() float32 {
	var numOfButtons = float32(len(software.InstalledSoftwareMap))
	var totalGroupWidth = ((SoftwareBtnWidth * numOfButtons) + (SoftwareBtnPadding * numOfButtons)) - SoftwareBtnPadding
	var posX = (float32(WindowW) / 2) - (totalGroupWidth / 2)
	return posX
}

// RenderActiveSoftwareButtons will generate the UI needed to display all found software installations
func RenderActiveSoftwareButtons() g.Widget {
	return g.Custom(func() {
		g.Button("Demo").Size(80, SoftwareBtnHeight).OnClick(func() {
			ShowDemoWindow = true
		}).Build()
		for i, softwareName := range software.AllActiveSoftwareNames {
			if _, ok := software.InstalledSoftwareMap[softwareName]; ok {
				if i == 0 {
					imgui.SameLineV(softwareBtnPosX(), 0)
				} else {
					imgui.SameLineV(softwareBtnPosX()+SoftwareBtnWidth+SoftwareBtnPadding, 0)
				}
				g.Button(softwareName).Size(SoftwareBtnWidth, SoftwareBtnHeight).OnClick(func() {
					ActiveSoftwareTab = softwareName
				}).Build()
			}
		}
		imgui.SameLineV(float32(WindowW)-SoftwareBtnHeight-float32(10), -1)
		g.Button("X").Size(SoftwareBtnHeight, SoftwareBtnHeight).OnClick(onQuit).Build()
	})
}

// RenderActiveSoftwareTab takes the active
func RenderActiveSoftwareTab() g.Widget {
	if _, ok := software.InstalledSoftwareMap[software.SoftwareVectorworks]; ok && ActiveSoftwareTab == ""{
		return g.Line(
			g.Custom(func() {
				installations, _ := software.FindInstallationsBySoftware(software.SoftwareVectorworks)
				for _, install := range installations {
					g.Label(install.Year).Build()
					imgui.SameLineV(0, 20)
					g.Label(install.License.Serial).Build()
				}
			}),
		)
	} else {
		return g.Line(
			g.Custom(func() {
				installations, _ := software.FindInstallationsBySoftware(ActiveSoftwareTab)
				for _, install := range installations {
					g.Label(install.Year).Build()
					imgui.SameLineV(0, 20)
					g.Label(install.License.Serial).Build()
				}
			}),
		)
	}
	return nil
}

// Quit
func onQuit() {
	os.Exit(0)
}

//type installations []software.Installation

//func (installations installations) toInterfaceSlice() []interface{} {
//	widgetInterface := make([]interface{}, len(installations))
//	for i := range installations {
//		widgetInterface[i] = installations[i]
//	}
//	return widgetInterface
//}


// NAVIGATE MAP
//func getSoftwareWidgetLabels() []g.Widget {
//	var labels []g.Widget
//
//	for key, _ := range software.InstalledSoftwareMap {
//		label := g.Name(key)
//		labels = append(labels, label)
//	}
//	return labels
//}
