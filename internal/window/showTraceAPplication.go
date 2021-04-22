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

func RenderTraceApplication() g.Widget {
	return g.Custom(func() {
		if ShowTraceApplication {
			var TargetFile string
			var stdOutPut bytes.Buffer
			g.Line(
				g.Button("Load File...").Size(-1, 30).OnClick(func() {
					TargetFile, _ = dialog.File().Filter("Application: .exe, .app", "exe", "app").Filter("All Files:  .*", "*").Load()

					command := exec.Command(TargetFile)

					var stdBuffer bytes.Buffer
					mw := io.MultiWriter(os.Stdout, &stdBuffer)

					command.Stdout = mw
					command.Stderr = mw

					// Run the binary and stream the output.
					if err := command.Run(); err != nil {
						log.Panicln(err)
					}

					// Stream the log output to the terminal window
					stdOutPut = stdBuffer
				}),
			).Build()
			if TargetFile != "" {
				imgui.BeginChildV("TraceApplication", imgui.Vec2{X: float32(WindowSize.Width), Y: 400}, true, 0)
				imgui.Text(stdOutPut.String())
				imgui.EndChild()
			}
		}
	})
}
