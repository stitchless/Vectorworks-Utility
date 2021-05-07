package ui

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
	"github.com/sqweek/dialog"
	"log"
	"os"
	"os/exec"
)

var traceBuffer bytes.Buffer
var logBuffer bytes.Buffer
var readLogs bool = false
var trailTraceLogs bool

type vectorworksLogs struct {
	Ts       string `json:"ts"`
	LogLvl   int    `json:"log_lvl"`
	SN       string `json:"sn"`
	Session  string `json:"session"`
	VWVer    string `json:"vw_ver"`
	Platform string `json:"platform"`
	OSVer    string `json:"os_ver"`
	Cat      string `json:"cat"`
	Message  string `json:"message"`
	Type     string `json:"type"`
}

// RenderLogging will render all logging UI
func RenderLogging() g.Widget {
	return g.Custom(func() {
		if featureTraceApplication == currentFeature {
			g.Line(
				g.Button("Load File...").Size(-1, 30).OnClick(func() {

					targetFile, err := dialog.File().Filter("Application: .exe, .app", "exe", "app").Filter("All Files:  .*", "*").Load()
					if err != nil {
						log.Println(err)
					} else {
						traceBuffer.Reset()
						targetFile = confirmTargetFile(targetFile)
						go traceApplication(targetFile)
					}
				}),
			).Build()
			imgui.BeginTabBar("##softwareLogsTabBar")

			if imgui.BeginTabItem("Software Trace##softwareTraceTabItem") {
				imgui.BeginChildV("##traceTabWindow", imgui.Vec2{X: -1, Y: float32(WindowSize.Height - 80)}, true, 0)
				bufferString := traceBuffer.String()
				// 14 == InputTextFlagsReadOnly | 16 == InputTextFlagsNoUndoRedo || InputText.go
				imgui.InputTextMultilineV("##traces", &bufferString, imgui.Vec2{X: -1, Y: -1}, 1<<14|1<<16, nil)
				if imgui.Checkbox("Trail Logs", &trailTraceLogs) {
					imgui.SetScrollY(imgui.ScrollMaxY())
				}
				imgui.EndChild()
				imgui.EndTabItem()
			}

			if imgui.BeginTabItem("Software Logs##softwareLogsTabItem") {
				imgui.BeginChildV("##logTabWindow", imgui.Vec2{X: -1, Y: float32(WindowSize.Height - 80)}, true, 0)
				logBufferString := logBuffer.String()
				showLogs()
				// 14 == InputTextFlagsReadOnly | 16 == InputTextFlagsNoUndoRedo || InputText.go
				imgui.InputTextMultilineV("##showLogs", &logBufferString, imgui.Vec2{X: -1, Y: -1}, 1<<14|1<<16, nil)
				imgui.EndChild()
				imgui.EndTabItem()
			}

			imgui.EndTabBar()
		}
	})
}

// traceApplication takes a target path, and runs the target.  The stderr and stdout are then
// captured and passed to a package variable traceBuffer
func traceApplication(targetFile string) {

	cmd := exec.Command(targetFile)

	logger := log.New(&traceBuffer, "", log.Ldate|log.Ltime)

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
		traceBuffer.WriteString(errMessage)
	}

	if err := cmd.Wait(); err != nil {
		errMessage := "Lost connection with running application.  Please close and run again. \n" + err.Error() + "\n"
		traceBuffer.WriteString(errMessage)
	}
}

// TODO: Set up a ticker that will periodically check the logs files for any changes
// This will be done using the modified date file stat of the file, and compare it
// against a previous loop, or simply time from current time. (maybe updated in the past 30 seconds)
// If changes are found set as the active log and begin parsing.

// FIXME: Non-sent logs will be duplicated in the buffer when transferred to the sent log file.
// A possible solution may have something to do with a string
// "SentDataDivider ==============================" between jason formats
// ---
// This may be a non issue if logs are parsed against other known data such as a time stamp
// prior to being applied to a buffer
// Another solution is to read the sent file once and use a ticker the only the non-sent logs.
// This will still capture all logs without having to rely on parsing and comparing.  This is subject
// to timing issues where the logs can be sent before the loop is run again.
// https://github.com/radovskyb/watcher
// https://github.com/fsnotify/fsnotify
// showLogs currently shows all logs once for all software found (Vectorworks, and Vision)
func showLogs() {
	if !readLogs {
		readLogs = true
		for _, softwareName := range software.AllActiveSoftwareNames {
			// Data Structure:::Log File

			// Test for installations of active software prior to making a table
			if len(software.AllInstalledSoftwareMap[softwareName]) == 0 {
				return
			}
			for _, installation := range software.AllInstalledSoftwareMap[softwareName] {
				file, err := os.Open(installation.LogFile)
				if err != nil {
					errors.New("error: could not open log file" + installation.LogFile)
				}
				var obj vectorworksLogs

				scanner := bufio.NewScanner(file)
				scanner.Split(bufio.ScanLines)
				for scanner.Scan() {
					err = json.Unmarshal(scanner.Bytes(), &obj)
					if err != nil {
						errors.New("error: could not unmarshal json")
					}
					logBuffer.WriteString("session: " + obj.Session + " message: " + obj.Message + "\n")
				}
			}
		}
	}
}
