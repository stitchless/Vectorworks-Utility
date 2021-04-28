package ui

import (
	"bufio"
	"bytes"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"github.com/sqweek/dialog"
	"io"
	"log"
	"os/exec"
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

func runApplication(ch chan []byte, targetFile string) {
	cmd := exec.Command(targetFile)

	outReader, outWriter := io.Pipe()
	errReader, errWriter := io.Pipe()

	cmd.Stdout = outWriter
	cmd.Stderr = errWriter

	//mw := io.MultiWriter(outWriter, errWriter, &buffer)
	mr := io.MultiReader(outReader, errReader, &buffer)



	// Is the transfer the output of the application, through the channel.

	go func() {
		reader := bufio.NewReader(mr)

		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				break
			}
			ch <- line
			ch <- []byte{'\n'}
		}
	}()

	if err := cmd.Run(); err != nil {
		log.Panicln(err)
	}
}
