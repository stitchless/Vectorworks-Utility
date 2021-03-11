package software

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
	"html/template"
)

var tmpl *template.Template

func GenerateTemplates(templateFS embed.FS) {

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

	tmpl = template.Must(template.New("homepage.html.tmpl").Funcs(funcMap).ParseFS(templateFS, "**/*.html.tmpl"))
}


// HandleMessages handles messages
// Here as the middle mag between Javascript and Go
// Json format strings come in
// Json format strings go out in the form out payload
func HandleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	var s string
	if len(m.Payload) > 0 && m.Name != "" {
		// Confirm message json format
		if err = json.Unmarshal(m.Payload, &s); err != nil {
			payload = err.Error()
			return
		}
	}
	switch m.Name {
	case "software":
		// Software
		if payload, err = software(s); err != nil {
			payload = err.Error()
			return
		}
	case "editSerial":
		if payload, err = editSerial(s); err != nil {
			payload = err.Error()
			return
		}
	}


	return
}

type render struct {
	Html string `json:"html_string"`
}

type htmlValues struct {
	Title       string
	Preloader   bool
	Description string
	Softwares   []Software
}

func software(s string) (r render, err error) {
	fmt.Println("SOFTWARE: " + s)
	templateValues := htmlValues{
		Preloader:   false,
		Title:       "Welcome to the Vectorworks Utility Tool",
		Description: "This utility will allow you to make a variety of changes to Vectorworks, Vision, and Vectorworks Cloud Services Desktop App.  Simply select an action from the list below...",
		Softwares:   allSoftwares,
	}

	var tpl bytes.Buffer
	if err = tmpl.ExecuteTemplate(&tpl, "homePage", templateValues); err != nil {
		fmt.Println("Send Help...")
		return
	}
	r.Html = tpl.String()
	fmt.Println(tpl.String())
	return
}

func editSerial(serial string) (r render, err error){
	fmt.Println("TESTING: " + serial)
	return

}