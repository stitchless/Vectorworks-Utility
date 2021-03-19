package gui

import (
	g "github.com/AllenDang/giu"
	"github.com/jpeizer/Vectorworks-Utility/internal/software"
)

// TODO: Fix row to use input.
func BuildRows() []*g.RowWidget {
	for _, Installations := range software.InstallationsMap {
		rows := make([]*g.RowWidget, len(Installations))
		for _, installation := range Installations {
			rows[installation.ID] = g.Row(
				g.Label(installation.Year),
				g.Label(installation.License.Serial),
			)
		}
		return rows
	}
	return nil
}