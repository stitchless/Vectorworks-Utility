package ui

// windowDimensions is the width and height of the MasterWindow
type windowDimensions struct {
	Width int
	Height int
}
// WindowSize dictates the master windows initial oldWindow size
// vec2{int, int}
var WindowSize = windowDimensions{800, 600}