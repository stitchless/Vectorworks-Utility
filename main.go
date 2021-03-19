package main

import (
	g "github.com/AllenDang/giu"
	"github.com/jpeizer/Vectorworks-Utility/internal/gui"
	"os"
)

func loop() {
	g.SingleWindowWithMenuBar("Vectorworks App Utility").Layout(
		g.MenuBar().Layout(
			g.Menu("File").Layout(
				g.MenuItem("Save"),
			),
		),
		g.Group().Layout(
			gui.SoftwareLabels()...,
		),
		g.Label("Vectorworks, Inc. Application Utility").Font(&fontRoboto),
		g.Label("Hello world from giu"),
		g.Line(
			g.Button("Quit").OnClick(onQuit),
		),
	)
}

func main() {
	wnd := g.NewMasterWindow("Vectorworks App Utility", 800, 600, g.MasterWindowFlagsNotResizable, LoadFont)
	wnd.Run(loop)
}

func onQuit() {
	os.Exit(0)
}