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
	windowW int = 600
	windowH int = 800
)

var showDemoWindow = false

func loop() {
	g.SingleWindow("Vectorworks App Utility").Layout(
		g.Line(
			g.Custom(func() {
				var posX float32 = 150
				for softwareLabel, _ := range software.InstalledSoftwareMap {
					imgui.SameLineV(posX, -1)
					g.Button(softwareLabel).Size(150, 34).OnClick(func(){
						fmt.Println(softwareLabel)
					}).Build()
				}
			}),
		),
		g.Line(
			g.Button("Quit").OnClick(onQuit),
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
	wnd := g.NewMasterWindow("Vectorworks App Utility", windowH, windowW, g.MasterWindowFlagsNotResizable, LoadFont)
	wnd.Run(loop)
}

func onQuit() {
	os.Exit(0)
}