package ui

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
)

// Start by
func init() {
	err := software.GenerateInstalledSoftwareMap()
	if err != nil {
		fmt.Errorf("could not generate installation map for %s: %v", software.AllActiveSoftwareNames, err)
	}
}

// RenderShowSerials shows serials of found supported software
func RenderShowSerials() g.Widget {
	// Setup table flags
	const imguiTableFlags imgui.TableFlags = imgui.TableFlags_SizingFixedFit |
		imgui.TableFlags_RowBg |
		imgui.TableFlags_ScrollX |
		imgui.TableFlags_BordersOuterH |
		imgui.TableFlags_BordersOuterV

	return g.Custom(func() {
		if featureShowSerial == currentFeature {
			// Run for all active supported software
			for _, softwareName := range software.AllActiveSoftwareNames {
				// Test for installations of active software prior to making a table
				if len(software.AllInstalledSoftwareMap[softwareName]) != 0 {
					// Spacing before the table entry
					imgui.Dummy(imgui.Vec2{X: -1, Y: 30})
					imgui.BeginGroup()
					imgui.Text(softwareName)
					// Begin of software table
					imgui.BeginTable(softwareName+"Table", 2, imguiTableFlags, imgui.Vec2{X: -1, Y: 150}, 0)
					imgui.Selectable(softwareName + "##Row")
					imgui.TableNextRow(0, 35)

					// Row Content
					for _, installation := range software.AllInstalledSoftwareMap[softwareName] {
						imgui.TableNextColumn()
						imgui.Text(installation.Year)
						imgui.TableNextColumn()
						g.Button(installation.License.Serial+"##"+softwareName).OnClick(func() {
							doSomething(softwareName, installation.ID)
						}).Build()
						//imgui.ButtonV(installation.License.Serial+"##"+softwareName, imgui.Vec2{X: -1, Y: 30})
						//imgui.Text(installation.License.Serial)
					}

					imgui.EndTable()
					imgui.EndGroup()
				}
			}
		}
	})
}

func doSomething(softwareName software.SoftwareName, ID int) {
	installations := software.AllInstalledSoftwareMap[softwareName]
	fmt.Println(installations[ID].License.Serial)
}
