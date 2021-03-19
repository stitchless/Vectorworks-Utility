package main

import (
	_ "embed"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
)

var (
	fontRoboto imgui.Font
	//go:embed assets/fonts/DancingScript-Regular.ttf
	dancingScript []byte
	//go:embed assets/fonts/Roboto-Regular.ttf
	roboto []byte
)

func LoadFont() {
	// Load Fonts
	fonts := g.Context.IO().Fonts()
	fonts.AddFontFromMemoryTTF(dancingScript, 18)
	//fonts.AddFontFromFileTTF("./assets/fonts/DancingScript-Regular.ttf", 18)
	fontRoboto = fonts.AddFontFromMemoryTTF(roboto, 18)
}
