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

// RenderShowSoftware shows serials of found supported software
func RenderShowSoftware() g.Widget {
	// Setup table flags
	const imguiTableFlags imgui.TableFlags = imgui.TableFlags_SizingFixedFit |
		imgui.TableFlags_RowBg |
		imgui.TableFlags_ScrollX |
		imgui.TableFlags_BordersOuterH |
		imgui.TableFlags_BordersOuterV

	return g.Custom(func() {
		if featureSoftware == currentFeature {
			// Run for all active supported software
			for _, softwareName := range software.AllActiveSoftwareNames {
				// Test for installations of active software prior to making a table
				if len(software.AllInstalledSoftwareMap[softwareName]) != 0 {
					// Spacing before the table entry
					imgui.Dummy(imgui.Vec2{X: -1, Y: 25})
					imgui.BeginGroup()
					imgui.Text(softwareName)
					// Begin of software table
					imgui.BeginTable(softwareName+"Table", 3, imguiTableFlags, imgui.Vec2{X: -1, Y: 150}, 0)
					imgui.TableNextRow(0, 25)

					// Row Content
					for _, installation := range software.AllInstalledSoftwareMap[softwareName] {
						imgui.TableNextColumn()
						imgui.Text(installation.Year)
						imgui.TableNextColumn()
						g.Button(installation.License.Serial + "##" + softwareName).OnClick(func() {
							doSomething(softwareName, installation.ID)
						}).Build()
						imgui.TableNextColumn()
						//serialTags := parseSerial(installation.License.Serial)
						// Inner table to show all available serial tags
						imgui.BeginTable("##"+softwareName+"Tags", 4, imgui.TableFlags_SizingFixedFit, imgui.Vec2{X: 500, Y: -1}, 0)
						imgui.TableNextRow(0, 25)
						//for _, tag := range serialTags {
						//	imgui.TableNextColumn()
						//	imgui.Selectable(tag)
						//}
						imgui.TableNextColumn()
						imgui.Selectable(installation.License.Platform)
						imgui.TableNextColumn()
						imgui.Selectable(installation.License.Type)
						imgui.TableNextColumn()
						imgui.Selectable(installation.License.Activation)
						imgui.TableNextColumn()
						imgui.Selectable(installation.License.Local)
						imgui.EndTable()
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
	fmt.Println(installations[ID].Properties)
}
