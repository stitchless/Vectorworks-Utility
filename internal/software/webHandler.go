package software

import (
	"html/template"
	"net/http"
)

type htmlValues struct {
	Title       string
	Preloader   bool
	Description string
	Softwares   []Software
	FormData    FormData
}

type FormData struct {
	Name   string
	Year   string
	Serial string
}

// TODO: Determine if VCS is installed

var tmpl *template.Template

// homePageHandler is the initial page with all software information held on it.
// Each time an action is done the user is returned to this screen
// From this screen you can edit license info or clean up application data.
func HomePageHandler(w http.ResponseWriter, _ *http.Request) {
	templateValues := htmlValues{
		Preloader:   false,
		Title:       "Welcome to the Vectorworks Utility Tool",
		Description: "This utility will allow you to make a variety of changes to Vectorworks, Vision, and Vectorworks Cloud Services Desktop App.  Simply select an action from the list below...",
		Softwares:   allSoftwares,
	}
	err := tmpl.ExecuteTemplate(w, "homePage", templateValues)
	Check(err)
}

// TODO: Show localizations via Tabs
// TODO: Show Actions as Modals? (No)
// TODO: Illustrate license types (Private Repo)
// editSerialHandler The screen to chose the user a text field to update a selected serial number
func EditSerialHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	Check(err)
	var softwareName string
	var appYear string
	var serial string

	for key, value := range r.Form {
		switch key {
		case "softwareName":
			softwareName = value[0]
		case "appYear":
			appYear = value[0]
		case "serial":
			serial = value[0]
		}
	}

	// Serve the screen
	templateValues := htmlValues{
		Preloader:   false,
		Title:       "Welcome to the Vectorworks Utility Tool",
		Description: "This utility will allow you to make a variety of changes to Vectorworks, Vision, and Vectorworks Cloud Services Desktop App.  Simply select an action from the list below...",
		Softwares:   allSoftwares,
		FormData: FormData{
			Name:   softwareName,
			Year:   appYear,
			Serial: serial,
		},
	}

	err = tmpl.ExecuteTemplate(w, "editSerial", templateValues)
	Check(err)
}

// updateSerialHandler will send the filled in text field and update the serial
// Once updated, the home homePageHandler is called and the home screen is shown
func UpdateSerialHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	var softwareName string
	var appYear string
	var serial string

	for key, value := range r.Form {
		switch key {
		case "softwareName":
			softwareName = value[0]
		case "appYear":
			appYear = value[0]
		case "serial":
			serial = value[0]
		}
	}

	installation := Installation{
		Software: softwareName,
		Year:     appYear,
	}

	replaceOldSerial(installation, serial)

	templateValues := htmlValues{
		Preloader:   false,
		Title:       "Welcome to the Vectorworks Utility Tool!",
		Description: "This utility will allow you to make a variety of changes to Vectorworks, Vision, and Vectorworks Cloud Services Desktop App.  Simply select an action from the list below...",
		Softwares:   allSoftwares,
		FormData: FormData{
			Name:   softwareName,
			Year:   appYear,
			Serial: serial,
		},
	}

	err = tmpl.ExecuteTemplate(w, "homePage", templateValues)
	Check(err)
}

// TODO: Add new method for showing the cleaning of the application.

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	Check(err)
	var softwareName string
	var appYear string

	for key, value := range r.Form {
		switch key {
		case "softwareName":
			softwareName = value[0]
		case "appYear":
			appYear = value[0]
		}
	}

	installations, _ := FindInstallationsBySoftware(softwareName)
	for _, installation := range installations {
		if installation.Year == appYear {
			installation.Clean()
		}
	}

	// Serve the screen
	templateValues := htmlValues{
		Preloader:   false,
		Title:       "Welcome to the Vectorworks Utility Tool",
		Description: "This utility will allow you to make a variety of changes to Vectorworks, Vision, and Vectorworks Cloud Services Desktop App.  Simply select an action from the list below...",
		Softwares:   allSoftwares,
	}

	err = tmpl.ExecuteTemplate(w, "editSerial", templateValues)
	Check(err)
}

func ClearUserFolder(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	Check(err)
	var softwareName string
	var softwareYear string

	for key, value := range r.Form {
		switch key {
		case "softwareName":
			softwareName = value[0]
		case "softwareYear":
			softwareYear = value[0]
		}
	}


	installations, _ := FindInstallationsBySoftware(softwareName)
	for _, installation := range installations {
		if installation.Year == softwareYear {
			installation.Clean()
		}
	}

	// Serve the screen
	templateValues := htmlValues{
		Preloader:   false,
		Title:       "Welcome to the Vectorworks Utility Tool",
		Description: "This utility will allow you to make a variety of changes to Vectorworks, Vision, and Vectorworks Cloud Services Desktop App.  Simply select an action from the list below...",
		Softwares:   allSoftwares,
	}

	err = tmpl.ExecuteTemplate(w, "editSerial", templateValues)
	Check(err)
}
