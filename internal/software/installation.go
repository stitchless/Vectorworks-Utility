package software

import "strings"

type Installation struct {
	ID int
	SoftwareName
	Year        string
	License     License
	Properties  []string
	RMCache     string
	LogFiles    []string
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
		serial := getSerial(installation)
		serialFirstGroup := strings.Split(serial[0:6], "")

		installation.setProperties()
		installation.setUserData()
		installation.setRMCache()
		installation.setLogFiles()
		installation.License = License{
			Serial:     serial,
			Local:      getLocal(serialFirstGroup),
			Platform:   getPlatform(serialFirstGroup),
			Activation: getActivation(serialFirstGroup),
			Type:       getType(serialFirstGroup),
		}

		installations = append(installations, installation)
		i += 1
	}

	return installations, nil
}

func getActivation(serial []string) string {
	out, OK := licenseActivationMap[serial[0]]
	if OK {
		return out
	}
	return "Activation not found"
}

func getPlatform(serial []string) string {
	out, OK := licensePlatformMap[serial[2]]
	if OK {
		return out
	}
	return "Platform not found"
}

func getLocal(serial []string) string {
	local := strings.Join(serial[3:5], "")
	out, OK := licenseLocalMap[local]
	if OK {
		return out
	}
	return "Local not found"
}

func getType(serial []string) string {
	out, OK := licenseTypeMap[serial[5]]
	if OK {
		return out
	}
	return "License type not found"
}
