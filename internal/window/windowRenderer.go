package window

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
	"os"
)

// softwareBtnPosX will calculate the middle of the window minus half the UI elements drawn to the line
func softwareBtnPosX() float32 {
	var numOfButtons = float32(len(software.InstalledSoftwareMap))
	var totalGroupWidth = ((BtnWidth * numOfButtons) + (BtnPadding * numOfButtons)) - BtnPadding
	var posX = (float32(WindowSize.Width) / 2) - (totalGroupWidth / 2)
	return posX
}

func ShowInstalledSoftware(name software.SoftwareName) g.Widget {
	return g.Line(
		g.Custom(func() {
			fmt.Println(name)
		}),
	)
}

func onQuit() {
	os.Exit(0)
}