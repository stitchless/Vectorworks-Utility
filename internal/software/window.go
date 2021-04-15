package software

const (
	BtnWidth   float32 = 150
	BtnHeight  float32 = 34
	BtnPadding float32 = 10
)

type WindowDimensions struct {
	Width int
	Height int
}

var WindowSize = WindowDimensions{800, 600}

var ShowDemoWindow = false
var ActiveSoftwareTab string