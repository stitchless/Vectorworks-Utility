package main

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"github.com/jpeizer/Vectorworks-Utility/internal/gui"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
	"os"
)

var softwares []interface{}

func loop() {
	g.SingleWindow("Vectorworks App Utility").Layout(
		g.Group().Layout(
			g.Line(
				g.RangeBuilder("Installed Software", softwares, func(i int, v interface{}) g.Widget {
					str := v.(string)
					return g.Layout{
						g.Label(str),
						g.Table("Fast table").FastMode(true).Rows(gui.BuildRows(str)...),
					}
				})...,
			),
		),
		g.Line(
			g.Button("Quit").OnClick(onQuit),
		),
	)
}

func main() {
	err := software.GenerateSoftwareMap()
	if err != nil {
		fmt.Println(err)
	}
	softwares = gui.GetSoftwareNames().ToInterfaceSlice()
	wnd := g.NewMasterWindow("Vectorworks App Utility", 800, 600, g.MasterWindowFlagsNotResizable, LoadFont)
	wnd.Run(loop)
}

func onQuit() {
	os.Exit(0)
}