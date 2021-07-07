package software

import (
	"errors"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)


// TODO: Replace all home directories with GetConfigDirectory
func FindInstallationYears(softwareName SoftwareName) ([]string, error) {
	var appdataFolder string
	var years []string

	// Different software has different locations
	switch softwareName {
	case SoftwareVectorworks:
		appdataFolder = os.Getenv("APPDATA") + "/Nemetschek/Vectorworks"
	case SoftwareVision:
		appdataFolder = os.Getenv("APPDATA") + "/Vision"
	default:
		return nil, errors.New("info: APPDATA not found for provided paths")
	}

	folders, err := ioutil.ReadDir(appdataFolder)
	if err != nil {
		return nil, errors.New("error: could not read the appdata folder")
	}

	for _, f := range folders {
		year := regexp.MustCompile("[0-9]+").FindString(f.Name())
		if year != "" {
			years = append(years, year)
		}
	}

	return years, nil
}

// setProperties will take in an installation and assign it's properties strings
func (installation *Installation) setProperties() {
	version := convertYearToVersion(installation.Year)

	switch installation.SoftwareName {
	case SoftwareVectorworks:
		installation.Properties = []string{
			"SOFTWARE\\Nemetschek\\Vectorworks " + version,
			"SOFTWARE\\VectorWorks",
		}
	case SoftwareVision:
		installation.Properties = []string{
			"ESP Vision",
			"SOFTWARE\\VectorWorks\\Vision " + installation.Year,
			"SOFTWARE\\VWVision\\Vision" + installation.Year,
		}
	}
}

// setUserData well set all user data based on the target software
func (installation *Installation) setUserData() {
	winAppData := os.Getenv("APPDATA") + "\\"
	winLocalAppData := os.Getenv("LOCALAPPDATA") + "\\"

	// Set Directories based on software found
	switch installation.SoftwareName {
	case SoftwareVectorworks:
		installation.Directories = []string{
			winAppData + installation.SoftwareName + "\\" + installation.Year,
			winAppData + installation.SoftwareName + " " + installation.Year + " Installer",
			winAppData + installation.SoftwareName + " " + installation.Year + " Updater",
			winAppData + "Nemetschek\\" + installation.SoftwareName + "\\" + installation.Year,
			winAppData + "Nemetschek\\" + installation.SoftwareName + "\\accounts",
			winAppData + "Nemetschek\\" + installation.SoftwareName + " Web Cache",
			winAppData + "vectorworks-installer",
			winAppData + "vectorworks-updater",
			winAppData + "vectorworks-updater-updater",
			winLocalAppData + "vectorworks-updater-updater",
			winLocalAppData + "Nemetschek",
		}
	case SoftwareVision:
		installation.Directories = []string{
			filepath.Join(winAppData, installation.SoftwareName, installation.Year),
			filepath.Join(winLocalAppData, "VisionUpdater"),
		}
	case SoftwareCloudServices:
		installation.Directories = []string{
			winAppData + "vectorworks-cloud-services-beta",
			winAppData + "vectorworks-cloud-services",
			winLocalAppData + "vectorworks-cloud-services-beta-updater",
		}
	}
}

// setRMCache sets the system path for the resource manager cache directory
func (installation *Installation) setRMCache() {
	winAppData := os.Getenv("APPDATA") + "\\"
	installation.RMCache = winAppData + "Nemetschek\\" + installation.SoftwareName + " RMCache\\rm" + installation.Year
}

// setLogFiles sets all the log files paths for the target software
func (installation *Installation) setLogFileSent() {
	winAppData := os.Getenv("APPDATA") + "\\"
	installation.LogFileSent = filepath.Join(winAppData, "Nemetschek", installation.SoftwareName, installation.Year, "VW User Log Sent.txt")
}

// setLogFiles sets all the log files paths for the target software
func (installation *Installation) setLogFile() {
	winAppData := os.Getenv("APPDATA") + "\\"
	installation.LogFile = filepath.Join(winAppData, "Nemetschek", installation.SoftwareName, installation.Year, "VW User Log.txt")
}

// Clean removes the registry entry for the target software
func (installation Installation) Clean() {
	fmt.Println("Hello")
	for _, property := range installation.Properties {
		k, err := registry.OpenKey(registry.CURRENT_USER, property, registry.ALL_ACCESS)
		Check(err)

		names, _ := k.ReadSubKeyNames(-1)

		for _, name := range names {
			_ = registry.DeleteKey(k, name)
		}
		_ = registry.DeleteKey(k, "")

		func(k registry.Key) {
			err = k.Close()
			if err != nil {
			}
		}(k)
	}
	for _, directory := range installation.Directories {
		err := os.RemoveAll(directory)
		if err != nil {
			errors.New("error: could not delete directory: " + directory)
		}
	}
}