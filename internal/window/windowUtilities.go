package window

import (
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
)

func RenderInstallations(installations []software.Installation) g.Widget {
	return g.Custom(func() {
		// Show a tab bar
		imgui.BeginTabBar("SoftwareYears")
		for _, installation := range installations {
			var test = imgui.BeginTabItem(installation.Year)
			if test {
				imgui.EndTabItem()
			}
		}
		imgui.EndTabBar()
		// Show a table
		imgui.BeginTable("Installations", 2, imgui.TableFlags_SizingStretchSame, imgui.Vec2{X: 330, Y: 50}, 0)
		imgui.TableSetupColumn("Year", imgui.TableColumnFlags_IndentEnable, 80, 0)
		imgui.TableSetupColumn("Serial", imgui.TableColumnFlags_IndentEnable, 250, 0)
		imgui.TableHeadersRow()
		for _, installation := range installations {
			imgui.TableNextRow(0, 30)
			// Year Column
			imgui.TableNextColumn()
			imgui.Text(installation.Year)
			// Serial Column
			imgui.TableNextColumn()
			imgui.Text(installation.License.Serial)
		}
		imgui.EndTable()
		//rows[0].BgColor(&(color.RGBA{R: 200, G: 100, B: 100, A: 255}))
	})
}