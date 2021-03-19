package main

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"github.com/jpeizer/Vectorworks-Utility/internal/gui"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
	"os"
)

func loop() {
	g.SingleWindow("Vectorworks App Utility").Layout(
		g.Group().Layout(
			gui.SoftwareLabels()...,
		),
		g.Line(
			g.RangeBuilder("Installed Software", []interface{}{fmt.Sprintf("%#q", gui.GetSoftwareNames())}, func(i int, v interface{}) g.Widget {
				str := v.(string)
				return g.Layout{
					g.Label(str),
					g.Table("Fast table").FastMode(true).Rows(gui.BuildRows(str)...),
				}
			})...,
		),
		g.Line(
			g.Button("Quit").OnClick(onQuit),
		),
	)
}

func main() {
	err := software.PopulateInstallations()
	if err != nil {
		fmt.Println(err)
	}
	wnd := g.NewMasterWindow("Vectorworks App Utility", 800, 600, g.MasterWindowFlagsNotResizable, LoadFont)
	wnd.Run(loop)
}

func onQuit() {
	os.Exit(0)
}