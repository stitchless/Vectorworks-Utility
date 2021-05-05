package ui

import (
	"bytes"
	"fmt"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
	"github.com/sqweek/dialog"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

var buffer bytes.Buffer
var logbuffer bytes.Buffer
var readLogs bool = false

func RenderTraceApplication() g.Widget {
	return g.Custom(func() {
		if featureTraceApplication == currentFeature {
			g.Line(
				g.Button("Load File...").Size(-1, 30).OnClick(func() {

					targetFile, err := dialog.File().Filter("Application: .exe, .app", "exe", "app").Filter("All Files:  .*", "*").Load()
					if err != nil {
						log.Println(err)
					} else {
						buffer.Reset()
						targetFile = confirmTargetFile(targetFile)
						go runApplication(targetFile)
					}
				}),
			).Build()
			imgui.BeginTabBar("##softwareLogsTabBar")

			if imgui.BeginTabItem("Software Trace##softwareTraceTabItem") {
				imgui.BeginChildV("##traceTabWindow", imgui.Vec2{X: -1, Y: float32(WindowSize.Height - 120)}, true, 0)
				bufferString := buffer.String()
				// 14 == InputTextFlagsReadOnly | 16 == InputTextFlagsNoUndoRedo || InputText.go
				imgui.InputTextMultilineV("##traces", &bufferString, imgui.Vec2{X: -1, Y: -1}, 1<<14|1<<16, nil)
				imgui.EndChild()
				imgui.EndTabItem()
			}

			if imgui.BeginTabItem("Software Logs##softwareLogsTabItem") {
				imgui.BeginChildV("##logTabWindow", imgui.Vec2{X: -1, Y: float32(WindowSize.Height - 120)}, true, 0)
				logBufferString := logbuffer.String()
				logs()
				// 14 == InputTextFlagsReadOnly | 16 == InputTextFlagsNoUndoRedo || InputText.go
				imgui.InputTextMultilineV("##logs", &logBufferString, imgui.Vec2{X: -1, Y: -1}, 1<<14|1<<16, nil)
				imgui.EndChild()
				imgui.EndTabItem()
			}

			imgui.EndTabBar()
		}
	})
}

// runApplication takes a target path, and runs the target.  The stderr and stdout are then captured and passed
// to a package variable buffer
func runApplication(targetFile string) {

	cmd := exec.Command(targetFile)

	logger := log.New(&buffer, "", log.Ldate|log.Ltime)

	logStreamerOut := NewLogstreamer(logger, "stdout", false)
	defer func(logStreamerOut *Logstreamer) {
		err := logStreamerOut.Close()
		if err != nil {
			fmt.Fprintln(os.Stdout, "Error with Stdout: ", err)
		}
	}(logStreamerOut)

	logStreamerErr := NewLogstreamer(logger, "stderr", true)
	defer func(logStreamerErr *Logstreamer) {
		err := logStreamerErr.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error with Stderr: ", err)
		}
	}(logStreamerErr)

	cmd.Stdout = logStreamerOut
	cmd.Stderr = logStreamerErr

	if err := cmd.Start(); err != nil {
		errMessage := "Error starting application, please check your settings and try again... \n" + err.Error() + "\n"
		buffer.WriteString(errMessage)
	}

	if err := cmd.Wait(); err != nil {
		errMessage := "Lost connection with running application.  Please close and run again. \n" + err.Error() + "\n"
		buffer.WriteString(errMessage)
	}
}

func logs() {
	if !readLogs {
		readLogs = true
		for _, softwareName := range software.AllActiveSoftwareNames {
			// Test for installations of active software prior to making a table
			if len(software.AllInstalledSoftwareMap[softwareName]) == 0 {
				return
			}
			for _, installation := range software.AllInstalledSoftwareMap[softwareName] {
				for _, logFile := range installation.LogFiles {
					lines := make(map[int]string)

					data, err := ioutil.ReadFile(logFile)
					if err != nil {
						fmt.Println("Cannot Open File")
					}

					for n, line := range strings.Split(string(data), "\n") {
						lines[n] = line
					}
					lastLine := len(lines) - 2
					fmt.Println(lines[lastLine])
					fmt.Println(len(lines))
					//err := json.Unmarshal(lines[2], )
					//if err != nil {
					//	return
					//}
				}
			}
		}
	}
}