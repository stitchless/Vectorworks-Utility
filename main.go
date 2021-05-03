package main

import (
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"github.com/jpeizer/Vectorworks-Utility/internal/ui"
)

func loop() {
	g.SingleWindow("Vectorworks Utility##MainWindow").Layout(
		// START CONTENT AREA
		// Static header for all available features
		g.Line(
			ui.RenderTopMenuBar(),
			g.Custom(func() {
				imgui.Separator()
			}),
		),
		// Feature Content
		ui.RenderTraceApplication(),
		ui.RenderShowSoftware(fontRobotoTitle, fontAwesome),
		ui.RenderDemoWindow(),
		// END CONTENT AREA
	)
}

func main() {
	// 0 == WindowFlagsNone
	var gMasterWindowFlags g.MasterWindowFlags = 0
	// imgui docs: https://github.com/AllenDang/imgui-go
	window := g.NewMasterWindow("Vectorworks Inc.", 1200, 850, gMasterWindowFlags, LoadFont)
	window.Run(loop)
}
