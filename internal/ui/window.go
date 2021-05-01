package ui

// windowDimensions is the width and height of the MasterWindow
type windowDimensions struct {
	Width int
	Height int
}
// WindowSize dictates the master windows initial oldWindow size
// vec2{int, int}
var WindowSize = windowDimensions{800, 600}









//// MasterWindowFlags Provides all available options for the Master Window
//type MasterWindowFlags struct {
//	noTitlebar	bool
//	noScrollbar bool
//	noMove bool
//	noResize bool
//	noCollapse bool
//	noBackground bool
//}
//
//func (f MasterWindowFlags) combined() g.WindowFlags {
//	flags := g.MasterWindowFlags(0)
//	if f.noTitlebar {
//		flags |= imgui.WindowFlagsNoTitleBar
//	}
//	if f.noScrollbar {
//		flags |= imgui.WindowFlagsNoScrollbar
//	}
//	if f.noMove {
//		flags |= imgui.WindowFlagsNoMove
//	}
//	if f.noResize {
//		flags |= imgui.WindowFlagsNoResize
//	}
//	if f.noCollapse {
//		flags |= imgui.WindowFlagsNoCollapse
//	}
//	if f.noBackground {
//		flags |= imgui.WindowFlagsNoBackground
//	}
//	return g.WindowFlags(flags)
//}