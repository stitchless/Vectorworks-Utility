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

var (
	toggleSerialDetails bool
)

//// Clipboard describes the access to the text clipboard of the window manager.
//type Clipboard struct {
//	content string
//}
//
//func (c *Clipboard) Text() (string, error) {
//	return c.content, nil
//}
//
//func (c *Clipboard) SetText(value string) {
//	c.content = value
//}
//
//var clipboard Clipboard

// RenderShowSoftware shows serials of found supported software
func RenderShowSoftware(fontRobotoTitle imgui.Font, fontAwesome imgui.Font) g.Widget {
	// Setup table flags
	const imguiTableFlags imgui.TableFlags = imgui.TableFlags_SizingFixedFit |
		imgui.TableFlags_RowBg |
		imgui.TableFlags_ScrollX |
		imgui.TableFlags_BordersOuterH |
		imgui.TableFlags_BordersOuterV

	return g.Custom(func() {
		if featureSoftware != currentFeature {
			return
		}
		// Start of software tab bar
		imgui.BeginTabBar("##SoftwareTabBar")
		// Run for all active supported software
		for _, softwareName := range software.AllActiveSoftwareNames {
			// Test for installations of active software prior to making a table
			if len(software.AllInstalledSoftwareMap[softwareName]) == 0 {
				continue
			}
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
						// Software serial label
						imgui.Dummy(imgui.Vec2{X: -1, Y: 5})
						imgui.PushFont(fontRobotoTitle)
						// Flags 2 InputTextFlagsCharsUppercase | 4 InputTextFlagsAutoSelectAll | InputTextFlagsEnterReturnsTrue
						if imgui.InputTextV("##EditedSerial", &installation.License.Serial, 1<<2|1<<4|1<<5, nil) {
							software.ReplaceOldSerial(installation, installation.License.Serial)
							err := software.GenerateInstalledSoftwareMap()
							if err != nil {
								fmt.Errorf("error updating internal installation data after serial update %v", err)
							}
						}
						imgui.PopFont()
						if imgui.IsItemHovered() {
							imgui.SetTooltip("Insert new serial and press enter to update")
						}


						// Cog Icon button
						imgui.SameLine()
						imgui.PushFont(fontAwesome)
						if imgui.Button("\uF013" + "##" + installation.Year + "licenseButton") {
							toggleSerialDetails = !toggleSerialDetails
						}
						imgui.PopFont()
						// Show License Tags
						if toggleSerialDetails {
							imgui.BeginTable("##softwareTagsTable", 4, imgui.TableFlags_SizingFixedFit, imgui.Vec2{X: -1, Y: 30}, 0)
							imgui.TableNextColumn()
							imgui.Text(installation.License.Platform)
							imgui.SameLine()
							imgui.Dummy(imgui.Vec2{X: 20, Y: -1})
							imgui.TableNextColumn()
							imgui.Text(installation.License.Local)
							imgui.SameLine()
							imgui.Dummy(imgui.Vec2{X: 20, Y: -1})
							imgui.TableNextColumn()
							imgui.Text(installation.License.Activation)
							imgui.SameLine()
							imgui.Dummy(imgui.Vec2{X: 20, Y: -1})
							imgui.TableNextColumn()
							imgui.Text(installation.License.Type)
							imgui.EndTable()
						}
						imgui.Dummy(imgui.Vec2{X: -1, Y: 5})
						imgui.BeginChildV("##softwareContentChild", imgui.Vec2{X: -1, Y: float32(WindowSize.Height - 120)}, true, 0)

						//////////
						// Edit Serial
						//////////
						imgui.Text("Testing")

						//////////
						// Clear User Data
						//////////

						imgui.EndChild()
						// ----------------------------
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
		// Ending the software name tab bar
		imgui.EndTabBar()
	})
}
