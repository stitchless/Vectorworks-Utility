package window

import (
	"bytes"
	"fmt"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	"github.com/go-cmd/cmd"
	"github.com/sqweek/dialog"
	"log"
	"time"
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
	streamCmd := cmd.NewCmd(targetFile)
	statusChan := streamCmd.Start()

	ticker := time.NewTimer(2 * time.Second)

	go func() {
		for range ticker.C {
			status := streamCmd.Status()
			n := len(status.Stdout)
			fmt.Println(status.Stdout[n])
		}
	}()

	select {
	case finalStatus := <-statusChan:
		// done
		fmt.Println(finalStatus)
	default:
		// no, still running
	}
	finalStatus := <-statusChan
	fmt.Println(finalStatus)

	//cmd := exec.Command(targetFile)
	//
	////var stdBuffer bytes.Buffer
	//stdoutPipe, _ := cmd.StdoutPipe()
	//stderrPipe, _ := cmd.StderrPipe()
	//multiReader := io.MultiReader(stdoutPipe, stderrPipe)
	//
	//go func() {
	//	bufioReader := bufio.NewReader(multiReader)
	//	for {
	//		output, _, err := bufioReader.ReadLine()
	//		var tempData []byte
	//		tempData = output
	//		if err != nil {
	//			break
	//		}
	//		ch <- tempData
	//		ch <- []byte{'\n'}
	//	}
	//}()
	//
	//if err := cmd.Run(); err != nil {
	//	log.Panicln(err)
	//}
}