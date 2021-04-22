package window// RenderActiveSoftwareTab takes the active
import (
	"fmt"
	g "github.com/AllenDang/giu"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
)

// RenderActiveSoftwareTab Find method for default view, Then follow click events
func RenderActiveSoftwareTab() g.Widget {
	if installations, ok := software.InstalledSoftwareMap[ActiveSoftwareTab]; ok {
		// Show the active software
		return g.Line(
			g.Custom(func() {
				RenderInstallations(installations).Build()
			}),
		)
	} else {
		// Provide default view if no selection is found
		return g.Line(
			g.Custom(func() {
				fmt.Println("Default Function")
				RenderInstallations(installations).Build()
			}),
		)
	}
}

