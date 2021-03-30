package main

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	"github.com/jpeizer/Vectorworks-Utility/internal/gui"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
	"os"
)

var softwares []interface{}

const (
	windowW int = 800
	windowH int = 600
	softwareBtnWidth float32 = 150
	softwareBtnHeight float32 = 34
	softwareBtnPadding float32 = 10
)

var showDemoWindow = false

type uiState struct {
	demoView	bool
	activeSoftwareTab string
}

func softwareBtnPosX() float32 {
	var numOfButtons = float32(len(software.InstalledSoftwareMap))
	var totalGroupWidth = ((softwareBtnWidth * numOfButtons) + (softwareBtnPadding * numOfButtons)) - softwareBtnPadding
	var posX = (float32(windowW) / 2) - (totalGroupWidth / 2)
	return posX
}

func loop() {
	g.SingleWindow("Vectorworks App Utility").Layout(
		g.Line(
			g.Custom(func() {
				for i, softwareName := range software.AllActiveSoftwareNames {
					if _, ok := software.InstalledSoftwareMap[softwareName]; ok {
						if i > 0 {
							imgui.SameLineV(softwareBtnPosX(), -1)
						} else {
							imgui.SameLineV(softwareBtnPosX() +softwareBtnPadding + softwareBtnWidth, -1)
						}
						g.Button(softwareName).Size(softwareBtnWidth, softwareBtnHeight).OnClick(func(){
							fmt.Println(softwareName)
							showDemoWindow = true
						}).Build()
					}
				}
				imgui.SameLineV(float32(windowW) - softwareBtnHeight - float32(10), -1)
				g.Button("X").Size(softwareBtnHeight, softwareBtnHeight).OnClick(onQuit).Build()
			}),
		),
	)

	if showDemoWindow {
		imgui.ShowDemoWindow(&showDemoWindow)
	}
}

func main() {
	err := software.GenerateInstalledSoftwareMap()
	if err != nil {
		fmt.Println(err)
	}
	softwares = gui.GetSoftwareNames().ToInterfaceSlice()
	wnd := g.NewMasterWindow(
		"Vectorworks App Utility",
		windowW,
		windowH,
		g.MasterWindowFlagsNotResizable,
		LoadFont,
	)
	wnd.Run(loop)
}

func onQuit() {
	os.Exit(0)
}