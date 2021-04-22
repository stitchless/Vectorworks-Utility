package window

import (
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
)

func RenderDemoWindow() g.Widget {
 return g.Custom(func() {
 	if ShowDemoWindow {
		imgui.ShowDemoWindow(&ShowDemoWindow)
	}
 })
}