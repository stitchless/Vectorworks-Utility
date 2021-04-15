package software

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	"os"
)

// softwareBtnPosX will calculate the middle of the window minus half the UI elements drawn to the line
func softwareBtnPosX() float32 {
	var numOfButtons = float32(len(InstalledSoftwareMap))
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
		for i, softwareName := range AllActiveSoftwareNames {
			if _, ok := InstalledSoftwareMap[softwareName]; ok {
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
	if _, ok := InstalledSoftwareMap[SoftwareVectorworks]; ok && ActiveSoftwareTab == ""{
		return g.Line(
			g.Custom(func() {
				installations, _ := FindInstallationsBySoftware(SoftwareVectorworks)
				fmt.Println("Default...")
				RenderInstallations(installations).Build()
				//g.Table("Test").Rows(RenderInstallations(installations)...).Size(500, 400).Build()
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
				installations, _ := FindInstallationsBySoftware(ActiveSoftwareTab)
				fmt.Println("Normal Function")
				RenderInstallations(installations)
				//g.Table("Test").FastMode(true).Rows(RenderInstallations(installations)...)
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