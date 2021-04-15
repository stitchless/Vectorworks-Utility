package software

const (
	SoftwareBtnWidth   float32 = 150
	SoftwareBtnHeight  float32 = 34
	SoftwareBtnPadding float32 = 10
)

type WindowDimensions struct {
	Width int
	Height int
}

var WindowSize = WindowDimensions{800, 600}

var ShowDemoWindow = false
var ActiveSoftwareTab string