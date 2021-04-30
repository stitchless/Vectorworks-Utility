package ui

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
	"strings"
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
						g.Button(installation.License.Serial+"##"+softwareName).OnClick(func() {
							doSomething(softwareName, installation.ID)
						}).Build()
						imgui.TableNextColumn()
						serialTags := parseSerial(installation.License.Serial)
						imgui.BeginTable("##"+softwareName+"Tags", 4, imgui.TableFlags_SizingFixedFit, imgui.Vec2{X: 500, Y: -1}, 0)
						imgui.TableNextRow(0, 25)
						for _, tag := range serialTags {
							imgui.TableNextColumn()
							imgui.Selectable(tag)
							//imgui.Button(tag)
						}
						imgui.EndTable()
						//imgui.Text(serialTags)
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
	fmt.Println(installations[ID].Properties)
}

func parseSerial(serial string) [4]string {
	parsedString := strings.Split(serial[0:6], "")
	var serialTags [4]string
	switch parsedString[0] {
	case "A":
		serialTags[0] = "Retired"
	case "B":
		serialTags[0] = "Local Dongle"
	case "C":
		serialTags[0] = "Network Dongle"
	case "E":
		serialTags[0] = "Serial Activation"
	case "G":
		serialTags[0] = "Network Server"
	case "U":
		serialTags[0] = "Updater"
	case "V":
		serialTags[0] = "Viewer"
	default:
		serialTags[0] = ""
	}

	switch parsedString[2] {
	case "X":
		serialTags[1] = "Cross Platform"
	case "M":
		serialTags[1] = "Mac"
	case "W":
		serialTags[1] = "Windows"
	}

	local := strings.Join(parsedString[3:5], "")
	switch local {
	case "US":
		serialTags[2] = "United States"
	case "MK":
		serialTags[2] = "United Kingdom"
	case "NZ":
		serialTags[2] = "New Zealand"
	case "ZC":
		serialTags[2] = "Australia"
	case "MR":
		serialTags[2] = "Canada (Resolve)"
	case "CA":
		serialTags[2] = "Canada (Paxar)"
	case "CL":
		serialTags[2] = "China"
	case "BZ":
		serialTags[2] = "Brazil"
	case "BE":
		serialTags[2] = "Belgium"
	default:
		serialTags[2] = ""
	}

	switch parsedString[5] {
	case "N":
		serialTags[3] = "Not For Resale (Retail)"
	case "E":
		serialTags[3] = "Educational (Pro Format)"
	case "S":
		serialTags[3] = "Student (Pro Format)"
	case "U":
		serialTags[3] = "Student (Student Format)"
	case "T":
		serialTags[3] = "Teacher (Pro Format)"
	case "C":
		serialTags[3] = "Teacher (Student Format)"
	case "A":
		serialTags[3] = "Internal"
	default:
		serialTags[3] = ""
	}
	return serialTags
}