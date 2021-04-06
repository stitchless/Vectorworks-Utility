package softwareWindow

import (
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
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

func RenderInstallations(installations []software.Installation) g.Widget {
	return g.Custom(func() {
		imgui.BeginTable("Installations", 1, imgui.TableFlags_SizingStretchSame, imgui.Vec2{X: 40, Y: 50}, 0)
		//g.Column("Testing").InnerWidthOrWeight(50)
		imgui.TableNextRow(0, 30)
		imgui.TableNextColumn()
		imgui.Text("TEST")
		//for i := range rows {
		//	rows[i] = g.Row(
		//		g.Label(fmt.Sprintf("%d", installations[i].ID)),
		//		g.Label(installations[i].License.Serial),
		//		g.Label(installations[i].SoftwareName),
		//		g.Column("Testing").InnerWidthOrWeight(50),
		//	)
		//}
		imgui.EndTable()
		//rows[0].BgColor(&(color.RGBA{R: 200, G: 100, B: 100, A: 255}))
	})
}