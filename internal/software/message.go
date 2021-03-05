package software

import (
	"encoding/json"
	"html/template"

	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

func GenerateTemplates() {
	// funcMap needed in order to define custom functions within go template
	funcMap := template.FuncMap{
		// Increments int by 1 (Used to illustrate table view)
		"inc": func(i int) int {
			return i + 1
		},
		"comp": func(val1 string, val2 string) bool {
			if val1 == val2 {
				return true
			}
			return false
		},
		"FindInstallationsBySoftware": FindInstallationsBySoftware,
	}

	// Gather templates and parse all found template files
	tmpl = template.Must(template.ParseGlob(GetWD() + "/web/template/*.html.tmpl")).Funcs(funcMap)
	template.Must(tmpl.ParseGlob(GetWD() + "/web/view/*.html.tmpl")).Funcs(funcMap)
}

// HandleMessages handles messages
func HandleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "home":
		// Unmarshal payload
		var path string
		if len(m.Payload) > 0 {
			// Unmarshal payload
			if err = json.Unmarshal(m.Payload, &path); err != nil {
				payload = err.Error()
				return
			}
		}

		// Explore
		//if payload, err = explore(path); err != nil {
		//	payload = err.Error()
		//	return
		//}
	}
	return
}
