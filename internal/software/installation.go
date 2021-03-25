package software

type Installation struct {
	ID          int
	License     License
	Properties  []string
	Directories []string
	Year        string
	SoftwareLabel
}

var InstalledSoftwareMap = make(map[string][]Installation)

func GenerateSoftwareMap() error {
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
