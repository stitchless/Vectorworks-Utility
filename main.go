package main

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
	"github.com/jpeizer/Vectorworks-Utility/internal/softwareWindow"
)

func loop() {
	g.SingleWindow("Vectorworks App Utility").Layout(
		g.Line(
			// Shows installed Software packages as buttons along with close button on the right
			softwareWindow.RenderActiveSoftwareButtons(),
		),
		g.Separator(),
		g.Line(
			g.Group().Layout(
				softwareWindow.RenderActiveSoftwareTab(),
			),
		),
	)

	if softwareWindow.ShowDemoWindow {
		imgui.ShowDemoWindow(&softwareWindow.ShowDemoWindow)
	}

	// Update window width for layout calculations
	var test = int(imgui.WindowWidth())
	fmt.Println(test)
}

func main() {
	err := software.GenerateInstalledSoftwareMap()
	software.Check(err)

	wnd := g.NewMasterWindow(
		"Vectorworks App Utility",
		softwareWindow.WindowW,
		softwareWindow.WindowH,
		0,
		LoadFont,
	)
	wnd.Run(loop)
}