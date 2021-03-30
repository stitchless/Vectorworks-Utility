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

func loop() {
	g.SingleWindow("Vectorworks App Utility").Layout(
		g.Custom(func() {
			imgui.Text("Aligned")
			imgui.SameLineV(150, -1)
			imgui.Text("Testing")
			imgui.SameLineV(300, -1)
			imgui.Text("Vision")
		}),
		g.Line(
			g.Button("Vectorworks").Size(150, 34).OnClick(func(){
				fmt.Println("Vectorworks Clicked")
			}),
			g.Button("Vision").Size(150, 34).OnClick(func(){
				fmt.Println("Vision Clicked")
			}),
		),
		//g.Group().Layout(
		//	g.Line(
		//		g.RangeBuilder("Installed Software", softwares, func(i int, v interface{}) g.Widget {
		//			str := v.(string)
		//			return g.Layout{
		//				g.Label(str),
		//				g.Table("Fast table").FastMode(true).Rows(gui.BuildRows(str)...),
		//			}
		//		})...,
		//	),
		//),
		//g.Line(
		//	g.Button("Quit").OnClick(onQuit),
		//),
	)
}

func main() {
	err := software.GenerateInstalledSoftwareMap()
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