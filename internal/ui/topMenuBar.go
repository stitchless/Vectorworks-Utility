package ui

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
)

// RenderTopMenuBar returns a widget containing a table with all active features of the application
func RenderTopMenuBar() g.Widget {
	numberOfButtons := len(AllActiveFeatures)
	return g.Custom(func() {
		// ID: 16 = ImGuiStyleVar_CellPadding on X value
		imgui.PushStyleVarVec2(16, imgui.Vec2{X: 8, Y: 0})
		imgui.BeginTable("Features##TopMenuBar", numberOfButtons, 0, imgui.Vec2{X: -1, Y: 30}, -1)
		imgui.TableNextRow(0, 30)
		for _, activeFeature := range AllActiveFeatures {
			if currentFeature == "" {
				currentFeature = activeFeature
			}
			imgui.TableNextColumn()
			g.Button(string(activeFeature)).Size(-1, 30).OnClick(func() {
				fmt.Println(currentFeature, activeFeature)
				currentFeature = activeFeature
			}).Build()
		}
		imgui.EndTable()
		imgui.PopStyleVar()
	})
}
