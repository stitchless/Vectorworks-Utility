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
			imgui.BeginChildV("showTraceApplication##traceWindow", imgui.Vec2{X: -1, Y: float32(WindowSize.Height - 120)}, true, 0)
			bufferString := buffer.String()
			// 14 == InputTextFlagsReadOnly | 16 == InputTextFlagsNoUndoRedo || InputText.go
			imgui.InputTextMultilineV("##traceLogs", &bufferString, imgui.Vec2{X: -1, Y: -1}, 1<<14|1<<16, nil)
			imgui.EndChild()
		}
	})
}

func runApplication(targetFile string) {
	// Idea from https://github.com/golang/go/issues/19685
	//
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
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		os.Exit(1)
	}
	<-done

	if err := cmd.Wait(); err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		os.Exit(1)
	}
}