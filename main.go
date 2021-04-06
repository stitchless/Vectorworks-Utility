package main

import (
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
	"github.com/jpeizer/Vectorworks-Utility/internal/softwareWindow"
)

var wnd *g.MasterWindow

func loop() {
	g.SingleWindow("Vectorworks App Utility").Layout(
		g.Line(
			// Shows installed Software packages as buttons along with close button on the right
			softwareWindow.RenderActiveSoftwareButtons(),
		),
		g.Separator(),
		g.Line(
			g.Group().Layout(
				// Display All installations for selected software
				softwareWindow.RenderActiveSoftwareTab(),
			),
		),
	)

	if softwareWindow.ShowDemoWindow {
		imgui.ShowDemoWindow(&softwareWindow.ShowDemoWindow)
	}

	softwareWindow.WindowSize.Width, _ = wnd.GetSize()
}

func main() {
	err := software.GenerateInstalledSoftwareMap()
	software.Check(err)

	wnd = g.NewMasterWindow(
		"Vectorworks App Utility",
		softwareWindow.WindowSize.Height,
		softwareWindow.WindowSize.Height,
		0,
		LoadFont,
	)
	wnd.Run(loop)
}