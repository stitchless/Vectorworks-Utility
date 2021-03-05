package software

type Installation struct {
	ID			int
	License     License
	Software    Software
	Properties  []string
	Directories []string
	Year        string
}

func FindInstallationsBySoftware(software Software) ([]Installation, error) {
	var installations []Installation

	years := FindInstallationYears(software)

	// Attach configs, versions, and Vectorworks Utility years all into on object then return that object
	for _, year := range years {
		installation := Installation{
			Software:    software,
			Year:        year,
		}

		installation.Properties = findProperties(installation)
		installation.Directories = findDirectories(installation)
		installation.License = License{
			Serial: getSerial(installation),
		}

		installations = append(installations, installation)
	}

	return installations, nil
}