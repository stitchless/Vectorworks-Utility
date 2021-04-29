package software

type Installation struct {
	ID          int
	SoftwareName
	Year        string
	License     License
	Properties  []string
	Directories []string
}

var AllInstalledSoftwareMap = make(map[SoftwareName][]Installation)

// GenerateInstalledSoftwareMap creates a map
// key: SoftwareName
// Value: [] Installation
func GenerateInstalledSoftwareMap() error {
	for _, softwareName := range AllActiveSoftwareNames {
		installations, err := FindInstallationsBySoftware(softwareName)
		if err != nil {
			return err
		}
		if len(installations) != 0 {
			AllInstalledSoftwareMap[softwareName] = installations
		}
	}
	return nil
}

// FindInstallationsBySoftware will take in a SoftwareName and build a slice of installs
// Returns a [] Installation
func FindInstallationsBySoftware(softwareLabel SoftwareName) ([]Installation, error) {
	var installations []Installation
	var i int

	years := FindInstallationYears(softwareLabel)

	// Attach configs, versions, and Vectorworks Utility years all into on object then return that object
	for _, year := range years {
		installation := Installation{
			ID:           i,
			Year:         year,
			SoftwareName: softwareLabel,
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
