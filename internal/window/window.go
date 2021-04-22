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

// WindowSize dictates the master windows initial window size
// vec2{int, int}
var WindowSize = windowDimensions{800, 600}

var ActiveSoftwareTab string

// Toggles all features during application run
var (
	ShowTraceApplication = false
	ShowEditSerial = false
	ShowDemoWindow = false
)