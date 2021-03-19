package software

type Installation struct {
	ID          int
	License     License
	Software    Software
	Properties  []string
	Directories []string
	Year        string
}

var InstallationsMap = make(map[string][]Installation)

func PopulateInstallations() error {
	for _, softwareName := range AllSoftwares {
		installations, err := FindInstallationsBySoftware(softwareName)
		if err != nil {
			return err
		}
		if len(installations) != 0 {
			InstallationsMap[softwareName] = installations
		}
	}
	return nil
}

func FindInstallationsBySoftware(software Software) ([]Installation, error) {
	var installations []Installation
	var i int

	years := FindInstallationYears(software)

	// Attach configs, versions, and Vectorworks Utility years all into on object then return that object
	for _, year := range years {
		installation := Installation{
			ID:       i,
			Software: software,
			Year:     year,
		}

		installation.Properties = findProperties(installation)
		installation.Directories = findDirectories(installation)
		installation.License = License{
			Serial: getSerial(installation),
		}

		installations = append(installations, installation)
		i += 1
	}

	return installations, nil
}
