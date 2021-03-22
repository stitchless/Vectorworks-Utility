package gui

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
)

// TODO: Fix row to use input softwarename as input
func BuildRows(softwareName string) []*g.RowWidget {
	rows := make([]*g.RowWidget, len(software.InstallationsMap[softwareName]))
	for i, installation := range software.InstallationsMap[softwareName] {
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