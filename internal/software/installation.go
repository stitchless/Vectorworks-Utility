package software

type Installation struct {
	ID          int
	License     License
	Properties  []string
	Directories []string
	Year        string
	SoftwareLabel
}

var InstalledSoftwareMap = make(map[SoftwareLabel][]Installation)

// GenerateInstalledSoftwareMap creates a map
// key: SoftwareLabel
// Value: [] Installation
func GenerateInstalledSoftwareMap() error {
	for _, softwareName := range AllSoftwareOptions {
		installations, err := FindInstallationsBySoftware(softwareName)
		if err != nil {
			return err
		}
		if len(installations) != 0 {
			InstalledSoftwareMap[softwareName] = installations
		}
	}
	return nil
}

// FindInstallationsBySoftware will take in a SoftwareLabel and build a slice of installs
// Returns a [] Installation
func FindInstallationsBySoftware(softwareLabel SoftwareLabel) ([]Installation, error) {
	var installations []Installation
	var i int

	years := FindInstallationYears(softwareLabel)

	// Attach configs, versions, and Vectorworks Utility years all into on object then return that object
	for _, year := range years {
		installation := Installation{
			ID:       i,
			Year:     year,
			SoftwareLabel: softwareLabel,
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
