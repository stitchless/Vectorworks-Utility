package main

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
	"github.com/jpeizer/Vectorworks-Utility/internal/uiElements"
)

func loop() {
	g.SingleWindow("Vectorworks App Utility").Layout(
		g.Line(
			// Shows installed Software packages as buttons along with close button on the right
			uiElements.RenderActiveSoftwareButtons(),
		),
		g.Separator(),
		g.Line(
			g.Group().Layout(
				uiElements.RenderActiveSoftwareTab(),
			),
		),
	)

	if uiElements.ShowDemoWindow {
		imgui.ShowDemoWindow(&uiElements.ShowDemoWindow)
	}
}

func main() {
	err := software.GenerateInstalledSoftwareMap()
	if err != nil {
		fmt.Println(err)
	}
	//softwares = uiElements.GetSoftwareNames().ToInterfaceSlice()
	wnd := g.NewMasterWindow(
		"Vectorworks App Utility",
		uiElements.WindowW,
		uiElements.WindowH,
		g.MasterWindowFlagsNotResizable,
		LoadFont,
	)
	wnd.Run(loop)
}
