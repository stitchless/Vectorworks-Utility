package main


// Embed and Load fonts for the UI to use


import (
	_ "embed"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
)

var (
	// Default Font
	//go:embed assets/fonts/nunito-sans-v6-latin-regular.ttf
	nunito []byte

	// FontRoboto Addressable Fonts
	FontRoboto imgui.Font
	//go:embed assets/fonts/Roboto-Regular.ttf
	roboto []byte
)

// LoadFont will load a default font then load addressable fonts that can be
// called for individual elements.
func LoadFont() {
	fonts := g.Context.IO().Fonts()
	fonts.AddFontFromMemoryTTF(nunito, 18)
	FontRoboto = fonts.AddFontFromMemoryTTF(roboto, 18)
}
