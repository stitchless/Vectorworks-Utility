package window

import (
	"bufio"
	"bytes"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
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
			case data := <- ch:
				buffer.Write(data)
			default:
		}

		if ShowTraceApplication {
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
			imgui.BeginChildV("TraceApplication", imgui.Vec2{X: float32(WindowSize.Width), Y: float32(WindowSize.Height - 120)}, true, 0)
			imgui.Text(buffer.String())
			imgui.EndChild()
			imgui.Button("Submit")
		}
	})
}

func runApplication(ch chan []byte, targetFile string) {
	command := exec.Command(targetFile)

	outReader, outWriter := io.Pipe()
	errReader, errWriter := io.Pipe()

	command.Stdout = outWriter
	command.Stderr = errWriter

	// Is the transfer the output of the application, through the channel.

	go func() {
		reader := bufio.NewReader(outReader)

		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				break
			}
			ch <- line
			ch <- []byte{'\n'}
		}
	}()

	go func() {
		reader := bufio.NewReader(errReader)

		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				break
			}
			ch <- line
			ch <- []byte{'\n'}
		}
	}()

	if err := command.Run(); err != nil {
		log.Panicln(err)
	}
}

