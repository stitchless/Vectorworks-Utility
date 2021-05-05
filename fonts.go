package main

// Embed and Load fonts for the UI to use
import (
	_ "embed"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
)

var (
	// TODO: Clean up this mess...
	// Default Font
	//go:embed assets/fonts/nunito-sans-v6-latin-regular.ttf
	nunito []byte

	// FontRoboto Addressable Fonts
	fontRoboto imgui.Font
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
	// Set up default fonts context
	fonts := g.Context.IO().Fonts()

	// Add font as default font for context
	fonts.AddFontFromMemoryTTF(roboto, 18)

	// Add FontAwesome on a per glyph basis
	ranges := imgui.NewGlyphRanges()
	builder := imgui.NewFontGlyphRangesBuilder()
	builder.AddText("\uF013")
	builder.BuildRanges(ranges)
	fontAwesome = fonts.AddFontFromMemoryTTFV(fa, 22, imgui.DefaultFontConfig, ranges.Data())

	// Add additional fonts that can be called by the application
	fontRobotoTitle = fonts.AddFontFromMemoryTTF(roboto, 22)
	fontRoboto = fonts.AddFontFromMemoryTTF(nunito, 18)
}
