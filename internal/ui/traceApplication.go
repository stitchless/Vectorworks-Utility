package ui

import (
	"bytes"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"github.com/sqweek/dialog"
	"log"
)

var buffer bytes.Buffer
var ch chan []byte

func init() {
	ch = make(chan []byte)
}

func RenderTraceApplication() g.Widget {
	return g.Custom(func() {
		select {
		case data := <-ch:
			buffer.Write(data)
		default:
		}

		if featureTraceApplication == currentFeature {
			g.Line(
				g.Button("Load File...").Size(-1, 30).OnClick(func() {

					targetFile, err := dialog.File().Filter("Application: .exe, .app", "exe", "app").Filter("All Files:  .*", "*").Load()
					if err != nil {
						log.Println(err)
					} else {
						go runApplication(ch, targetFile)
					}
				}),
			).Build()
			imgui.BeginChildV("showTraceApplication##TraceWindow", imgui.Vec2{X: -1, Y: float32(WindowSize.Height - 120)}, true, 0)
			imgui.Text(buffer.String())
			imgui.EndChild()
			imgui.Button("Submit")
		}
	})
}
