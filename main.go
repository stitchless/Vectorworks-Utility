package main

import (
	"fmt"
	"github.com/inkyblackness/imgui-go"
	"github.com/jpeizer/Vectorworks-Utility/internal/ui"
	"github.com/jpeizer/Vectorworks-Utility/internal/ui/platforms"
	"github.com/jpeizer/Vectorworks-Utility/internal/ui/renderers"
	"os"
	"time"
)

type clipboard struct {
	platform ui.Platform
}

func (board clipboard) Text() (string, error) {
	return board.platform.ClipboardText()
}

func (board clipboard) SetText(text string) {
	board.platform.SetClipboardText(text)
}

const (
	millisPerSecond = 1000
	sleepDuration   = time.Millisecond * 25
)

func main() {
	// https://github.com/inkyblackness/imgui-go-examples <- Documentation starting point is based on.
	context := imgui.CreateContext(nil)
	defer context.Destroy()
	io := imgui.CurrentIO()

	platform, err := platforms.NewGLFW(io, platforms.GLFWClientAPIOpenGL3)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer platform.Dispose()

	renderer, err := renderers.NewOpenGL3(io)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer renderer.Dispose()

	Run(platform, renderer)
}

func Run(p ui.Platform, r ui.Renderer) {
	imgui.CurrentIO().SetClipboard(clipboard{platform: p})

	clearColor := [3]float32{0.0, 0.0, 0.0}
	for !p.ShouldStop() {
		p.ProcessEvents()

		p.NewFrame()
		imgui.NewFrame()

		{
			imgui.Begin("MainWindow")
			imgui.Text("Oh Herro!")
			imgui.End()
		}

		imgui.Render()

		r.PreRender(clearColor)

		r.Render(p.DisplaySize(), p.FramebufferSize(), imgui.RenderedDrawData())
		p.PostRender()

		<-time.After(sleepDuration)
	}
}