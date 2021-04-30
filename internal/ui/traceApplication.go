package ui

import (
	"bufio"
	"bytes"
	"fmt"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"github.com/sqweek/dialog"
	"log"
	"os"
	"os/exec"
)

var buffer bytes.Buffer
var done chan struct{}

func init() {
	//done = make(chan struct{})
}

func RenderTraceApplication() g.Widget {
	return g.Custom(func() {
		if featureTraceApplication == currentFeature {
			g.Line(
				g.Button("Load File...").Size(-1, 30).OnClick(func() {

					targetFile, err := dialog.File().Filter("Application: .exe, .app", "exe", "app").Filter("All Files:  .*", "*").Load()
					if err != nil {
						log.Println(err)
					} else {
						targetFile = confirmTargetFile(targetFile)
						go runApplication(targetFile)
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

func runApplication(targetFile string) {
	// TODO: Check if application is still running
	// TODO: Close channel if application is no longer found
	cmd := exec.Command(targetFile)
	if cmdReader, err := cmd.StdoutPipe(); err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	} else {
		done = make(chan struct{})
		scanner := bufio.NewScanner(cmdReader)
		go func() {
			for scanner.Scan() {
				newLineByte := []byte("\n")
				logEntry := append(scanner.Bytes(), newLineByte...)
				buffer.Write(logEntry)
			}
			done <- struct{}{}
		}()
	}

	if err := cmd.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting Cmd", err)
		os.Exit(1)
	}
	<-done

	if err := cmd.Wait(); err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		os.Exit(1)
	}
}