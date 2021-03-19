package gui

import (
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
		)
	}
	//for i, installation := range rows {
	//	rows[i] = g.Row(
	//		g.Label(installation),
	//		g.Label(installation),
	//	)
	//}
	return rows
}