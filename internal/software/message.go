package software

import (
	"bytes"
	"embed"
	"encoding/json"
	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
	"html/template"
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
	//go:embed ../../resources/template/*
	var f embed.FS
	tmpl = template.Must(template.ParseFS(f, "**/*.html.tmpl")).Funcs(funcMap)
}

// HandleMessages handles messages
func HandleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "software":
		// Unmarshal payload
		var s string
		if len(m.Payload) > 0 {
			// Unmarshal payload
			if err = json.Unmarshal(m.Payload, &s); err != nil {
				payload = err.Error()
				return
			}
		}

		// Software
		if payload, err = software(); err != nil {
			payload = err.Error()
			return
		}
	}
	return
}

type Template struct {
	Html string `json:"html_string"`
}

func software() (t Template, err error) {
	templateValues := htmlValues{
		Preloader:   false,
		Title:       "Welcome to the Vectorworks Utility Tool",
		Description: "This utility will allow you to make a variety of changes to Vectorworks, Vision, and Vectorworks Cloud Services Desktop App.  Simply select an action from the list below...",
		Softwares:   allSoftwares,
	}

	var tpl bytes.Buffer
	if err = tmpl.ExecuteTemplate(&tpl, "homePage", templateValues); err != nil {
		return
	}
	t.Html = tpl.String()
	return
}