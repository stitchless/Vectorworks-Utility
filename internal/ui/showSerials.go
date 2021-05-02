package ui

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
)

// TODO: Move this to the software OnClick event for better initialization
func init() {
	err := software.GenerateInstalledSoftwareMap()
	if err != nil {
		fmt.Errorf("could not generate installation map for %s: %v", software.AllActiveSoftwareNames, err)
	}
}

// RenderShowSoftware shows serials of found supported software
func RenderShowSoftware(FontRobotoTitle imgui.Font) g.Widget {
	// Setup table flags
	const imguiTableFlags imgui.TableFlags = imgui.TableFlags_SizingFixedFit |
		imgui.TableFlags_RowBg |
		imgui.TableFlags_ScrollX |
		imgui.TableFlags_BordersOuterH |
		imgui.TableFlags_BordersOuterV

	return g.Custom(func() {
		if featureSoftware == currentFeature {
			// Start of software tab bar
			imgui.BeginTabBar("##SoftwareTabBar")
			// Run for all active supported software
			for _, softwareName := range software.AllActiveSoftwareNames {
				// Test for installations of active software prior to making a table
				if len(software.AllInstalledSoftwareMap[softwareName]) != 0 {
					// Insert new tab for each installed supported software
					if imgui.BeginTabItem(softwareName + "##" + softwareName + "TabItem") {
						// Begin of software year tab bar
						imgui.BeginTabBar("##" + softwareName + "TabBar")
						// Find all installed software versions
						for _, installation := range software.AllInstalledSoftwareMap[softwareName] {
							// Insert a new tab for all software versions found
							if imgui.BeginTabItem(installation.Year + "##" + softwareName + installation.Year + "TabItem") {
								// ----------------------------
								// LAYOUT FOR SOFTWARE FEATURES
								// ----------------------------
								
								imgui.Text(installation.License.Serial)
								// Ending the active software version tab content
								imgui.EndTabItem()
							}
						}
						// Ending the software version tab bar
						imgui.EndTabBar()
						// Ending the software name tab content
						imgui.EndTabItem()
					}
				}
			}
			// Ending the software name tab bar
			imgui.EndTabBar()
		}
	})
}

func doSomething(softwareName software.SoftwareName, ID int) {
	installations := software.AllInstalledSoftwareMap[softwareName]
	fmt.Println(installations[ID].Properties)
}
