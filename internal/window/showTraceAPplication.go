package window

import (
	"bytes"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	"github.com/sqweek/dialog"
	"io"
	"log"
	"os"
	"os/exec"
)

var stdOutPut bytes.Buffer

func RenderTraceApplication() g.Widget {
	return g.Custom(func() {
		if ShowTraceApplication {
			var TargetFile string
			g.Line(
				g.Button("Load File...").Size(-1, 30).OnClick(func() {
					var err error
					TargetFile, err = dialog.File().Filter("Application: .exe, .app", "exe", "app").Filter("All Files:  .*", "*").Load()
					if err != nil {
						log.Println(err)
					} else {
						go runApplication(TargetFile)
					}
				}),
			).Build()
			imgui.BeginChildV("TraceApplication", imgui.Vec2{X: float32(WindowSize.Width), Y: float32(WindowSize.Height-120)}, true, 0)
			imgui.Text(stdOutPut.String())
			imgui.EndChild()
			imgui.Button("Submit")
		}
	})
}

func runApplication(TargetFile string) {
	command := exec.Command(TargetFile)

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	command.Stdout = mw
	command.Stderr = mw

	// Run the binary and stream the output.
	if err := command.Run(); err != nil {
		log.Panicln(err)
	}

	// Stream the log output
	stdOutPut = stdBuffer
}