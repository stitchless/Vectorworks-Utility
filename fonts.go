package main

// Embed and Load fonts for the UI to use

//// #include "IconsFontAwesome5Pro.h"
//import "C"

import (
	_ "embed"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
)

var (
	// Default Font
	//go:embed assets/fonts/nunito-sans-v6-latin-regular.ttf
	nunito []byte

	// FontRoboto Addressable Fonts
	fontNunito imgui.Font
	//go:embed assets/fonts/Roboto-Regular.ttf
	roboto []byte

	fontRobotoTitle imgui.Font

	// FontRoboto Addressable Fonts
	fontAwesome imgui.Font
	//go:embed assets/fonts/fa-regular-400.ttf
	fa []byte
)

// LoadFont will load a default font then load addressable fonts that can be
// called for individual elements.
func LoadFont() {
	fonts := g.Context.IO().Fonts()
	fonts.AddFontFromMemoryTTF(roboto, 18)
	ranges := imgui.NewGlyphRanges()
	builder := imgui.NewFontGlyphRangesBuilder()
	builder.AddText("\uF013")
	builder.BuildRanges(ranges)
	fontAwesome = fonts.AddFontFromMemoryTTFV(fa, 18, imgui.DefaultFontConfig, ranges.Data())
	//fontAwesome = fonts.AddFontFromMemoryTTF(fa, 18)
	fontRobotoTitle = fonts.AddFontFromMemoryTTF(roboto, 22)
	//fontNunito = fonts.AddFontFromMemoryTTF(nunito, 18)
	//fontAwesome = fonts.AddFontFromMemoryTTF(fa, 18)

}
