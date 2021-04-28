package ui

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
)

// RenderDemoWindow shows the demo window as a reference while developing
func RenderDemoWindow() g.Widget {
	return g.Custom(func() {
		if featureDemoWindow == currentFeature {
			open := true
			imgui.ShowDemoWindow(&open)
			if !open {
				currentFeature = "DemoWindowClosed"
				fmt.Println("Entered")
			}
		}
	})
}
