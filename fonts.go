package main

// Embed and Load fonts for the UI to use
import (
	_ "embed"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
)

var (
	//go:embed assets/fonts/nunito-sans-v6-latin-regular.ttf
	nunito []byte
	//go:embed assets/fonts/Roboto-Regular.ttf
	roboto []byte
	//go:embed assets/fonts/fa-regular-400.ttf
	fa []byte

	FontAwesome imgui.Font
	FontRobotoTitle imgui.Font
	FontNunito imgui.Font
)

// LoadFont will load a default font then load addressable fonts that can be
// called for individual elements.
func LoadFont() {
	// Set up default fonts context
	fonts := g.Context.IO().Fonts()
	fonts.Clear()

	// Add font as default font for context
	fonts.AddFontFromMemoryTTF(roboto, 18)

	// Add FontAwesome on a per glyph basis
	ranges := imgui.NewGlyphRanges()
	builder := imgui.NewFontGlyphRangesBuilder()
	builder.AddText("\uF013")
	builder.BuildRanges(ranges)
	//fonts.AddFontFromMemoryTTFV(fa, 22, imgui.DefaultFontConfig, ranges.Data())
	fonts.AddFontFromMemoryTTFV(fa, 22, imgui.DefaultFontConfig, ranges.Data())

	// Add additional fonts that can be called by the application
	fonts.AddFontFromMemoryTTF(roboto, 22)
	FontNunito = fonts.AddFontFromMemoryTTF(nunito, 18)
	fonts.Build()
}
