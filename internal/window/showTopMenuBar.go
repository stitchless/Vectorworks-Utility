package window

import "C"
import (
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
)

// TODO: create new line
// TODO: Create group
// TODO: Show buttons to navigate the difference features

// RenderTopMenuBar CheckForActiveFeatures gets all active features and returns a slice of software.FeatureName
func RenderTopMenuBar() g.Widget {
	numberOfButtons := len(software.AllActiveFeatures)
	return g.Custom(func() {
		// ID = (16)ImGuiStyleVar_CellPadding on X value
		imgui.PushStyleVarVec2(16, imgui.Vec2{X: 8, Y: 0})
		imgui.BeginTable("TopMenuBar", numberOfButtons, 0, imgui.Vec2{X: -1, Y: 30}, 0)
		imgui.TableNextRow(0, 30)
		for _, feature := range software.AllActiveFeatures {
			imgui.TableNextColumn()
			g.Button(feature).Size(-1, 30).OnClick(func() {
				RunModule(feature)
			}).Build()
		}
		imgui.EndTable()
		imgui.PopStyleVar()
	})
}
