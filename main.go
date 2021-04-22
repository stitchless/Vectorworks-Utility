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
			window.RenderTopMenuBar(),
		),
		g.Separator(),
		g.Line(
			window.RenderTraceApplication(),
			// TODO: Fill this area with conditional content
			//window.RenderActiveSoftwareButtons(),
		),
	)

	if window.ShowDemoWindow {
		imgui.ShowDemoWindow(&window.ShowDemoWindow)
	}

	if window.ShowTraceApplication {
		window.RenderTraceApplication()
	}

	window.WindowSize.Width, window.WindowSize.Height = wnd.GetSize()
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