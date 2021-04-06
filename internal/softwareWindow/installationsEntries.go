package softwareWindow

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
	"image/color"
)

//func InstallationEntries() []g.Widget {
//	var widget []g.Widget
//	for _, Installations := range software.InstalledSoftwareMap {
//		for _, installation := range Installations {
//			year := g.Label(installation.Year)
//			serial := g.Label(installation.License.Serial)
//			widget = append(widget, year, serial)
//		}
//	}
//	return widget
//}

func RenderInstallations(installations []software.Installation) []*g.RowWidget {
	rows := make([]*g.RowWidget, len(installations))
	for i := range rows {
		rows[i] = g.Row(
			g.Label(fmt.Sprintf("%d", installations[i].ID)),
			g.Label(installations[i].License.Serial),
			g.Label(installations[i].SoftwareName),
		)
	}
	rows[0].BgColor(&(color.RGBA{R: 200, G: 100, B: 100, A: 255}))
	return rows
}