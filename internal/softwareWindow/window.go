package softwareWindow

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
	"os"
)

// softwareBtnPosX will calculate the middle of the window minus half the UI elements drawn to the line
func softwareBtnPosX() float32 {
	var numOfButtons = float32(len(software.InstalledSoftwareMap))
	var totalGroupWidth = ((SoftwareBtnWidth * numOfButtons) + (SoftwareBtnPadding * numOfButtons)) - SoftwareBtnPadding
	var posX = (float32(WindowSize.Width) / 2) - (totalGroupWidth / 2)
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
		imgui.SameLineV(float32(WindowSize.Width)-SoftwareBtnHeight-float32(10), -1)
		g.Button("X").Size(SoftwareBtnHeight, SoftwareBtnHeight).OnClick(onQuit).Build()
	})
}

// RenderActiveSoftwareTab takes the active
func RenderActiveSoftwareTab() g.Widget {
	if _, ok := software.InstalledSoftwareMap[software.SoftwareVectorworks]; ok && ActiveSoftwareTab == ""{
		return g.Line(
			g.Custom(func() {
				installations, _ := software.FindInstallationsBySoftware(software.SoftwareVectorworks)
				fmt.Println("Default...")
				g.Table("Test").FastMode(true).Rows(RenderInstallations(installations)...).Build()
				//for _, install := range installations {
				//	g.Label(install.Year).Build()
				//	imgui.SameLineV(0, 20)
				//	g.Label(install.License.Serial).Build()
				//}
			}),
		)
	} else {
		return g.Line(
			g.Custom(func() {
				installations, _ := software.FindInstallationsBySoftware(ActiveSoftwareTab)
				fmt.Println("Normal Function")
				g.Table("Test").FastMode(true).Rows(RenderInstallations(installations)...)
				//for _, install := range installations {
				//	g.Label(install.Year).Build()
				//	imgui.SameLineV(0, 20)
				//	g.Label(install.License.Serial).Build()
				//}
			}),
		)
	}
}

func onQuit() {
	os.Exit(0)
}