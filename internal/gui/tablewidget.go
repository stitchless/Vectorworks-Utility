package gui

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
)

func BuildRows(softwareLabel software.SoftwareName) []*g.RowWidget {
	rows := make([]*g.RowWidget, len(software.InstalledSoftwareMap[softwareLabel]))
	for i, installation := range software.InstalledSoftwareMap[softwareLabel] {
		rows[i] = g.Row(
			g.Label(installation.Year),
			g.Label(installation.License.Serial),
			g.Button("Edit Serial").OnClick(func(){
				fmt.Println("Edit Serial was Clicked...")
			}),
			g.Button("Delete User Folder"),
		)
	}
	return rows
}