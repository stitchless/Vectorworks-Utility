package softwareWindow

const (
	SoftwareBtnWidth   float32 = 150
	SoftwareBtnHeight  float32 = 34
	SoftwareBtnPadding float32 = 10
)

type WindowDimentions struct {
	Width int
	Height int
}

var WindowSize = WindowDimentions{800, 600}

var ShowDemoWindow = false
var ActiveSoftwareTab string


