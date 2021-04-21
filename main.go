package main

import (
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
	"github.com/jpeizer/Vectorworks-Utility/internal/window"
)

var wnd *g.MasterWindow

func loop() {
	g.SingleWindow("Vectorworks App Utility").Layout(
		g.Line(
			// Shows installed Software packages as buttons along with close button on the right
			window.RenderActiveSoftwareButtons(),
			window.ShowInstalledSoftware(window.ActiveSoftwareTab),
		),
		g.Separator(),
		g.Line(
			g.Group().Layout(
				// Display All installations for selected software
				window.RenderActiveSoftwareTab(),
			),
		),
	)

	if window.ShowDemoWindow {
		imgui.ShowDemoWindow(&window.ShowDemoWindow)
	}

	window.WindowSize.Width, _ = wnd.GetSize()
}

func main() {
	err := software.GenerateInstalledSoftwareMap()
	software.Check(err)

	wnd = g.NewMasterWindow(
		"Vectorworks App Utility",
		window.WindowSize.Height,
		window.WindowSize.Height,
		0,
		LoadFont,
	)
	wnd.Run(loop)
}