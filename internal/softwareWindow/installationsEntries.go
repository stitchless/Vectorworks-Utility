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
		imgui.BeginTable(
			"Installations",
			3,
			imgui.TableFlags_SizingStretchSame |
			imgui.TableFlags_ScrollY,
			imgui.Vec2{X: 400, Y: 50},
			0)
		imgui.TableSetupColumn("year", imgui.TableColumnFlags_WidthStretch | imgui.TableColumnFlags_IndentEnable, 100, 0)
		imgui.TableSetupColumn("Serial", imgui.TableColumnFlags_WidthStretch, 400, 0)
		imgui.TableSetupColumn("Copy", imgui.TableColumnFlags_WidthStretch, 70, 0)
		//g.Column("Testing").InnerWidthOrWeight(50)
		for i := range installations {
			imgui.TableNextRow(0, 30)
			imgui.TableNextColumn()
			imgui.Text(installations[i].Year)
			imgui.TableNextColumn()
			imgui.Text(installations[i].License.Serial)
			imgui.TableNextColumn()
			imgui.Button("Copy")
		}
		imgui.EndTable()
		//rows[0].BgColor(&(color.RGBA{R: 200, G: 100, B: 100, A: 255}))
	})
}