package window

const (
	BtnWidth   float32 = 150
	BtnHeight  float32 = 34
	BtnPadding float32 = 10
)

type windowDimensions struct {
	Width int
	Height int
}


var WindowSize = windowDimensions{800, 600}

var ShowDemoWindow = false
var ShowTraceApplication = false
var ActiveSoftwareTab string