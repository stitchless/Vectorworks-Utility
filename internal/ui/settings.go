package ui

import (
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
)

func ShowSettings() {
	g.PopupModal("Settings").Layout(
		g.Custom(func() {
			imgui.Text("Test")
		}),
	).Build()
}
