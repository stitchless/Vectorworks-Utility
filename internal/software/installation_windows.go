package software

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func FindInstallationYears(softwareLabel SoftwareName) []string {
	var appdataFolder string
	var years []string

	// Different software has different locations
	switch softwareLabel {
	case SoftwareVectorworks:
		appdataFolder = os.Getenv("APPDATA") + "/Nemetschek/Vectorworks"
	case SoftwareVision:
		appdataFolder = os.Getenv("APPDATA") + "/Vision"
	default:
		return nil
	}

	folders, _ := ioutil.ReadDir(appdataFolder)

	for _, f := range folders {
		year := regexp.MustCompile("[0-9]+").FindString(f.Name())
		if year != "" {
			years = append(years, year)
		}
	}

	return years
}

func findProperties(installation Installation) []string {
	// define system variables
	version := convertYearToVersion(installation.Year)

	switch installation.SoftwareName {
	case SoftwareVectorworks:
		return []string{
			"SOFTWARE\\Nemetschek\\Vectorworks " + version,
			"SOFTWARE\\VectorWorks",
		}
	case SoftwareVision:
		return []string{
			"ESP Vision",
			"SOFTWARE\\VectorWorks\\Vision " + installation.Year,
			"SOFTWARE\\VWVision\\Vision" + installation.Year,
		}
	}

	return nil
}

func findDirectories(installation Installation) []string {
	// define system variables
	winAppData := os.Getenv("APPDATA") + "\\"
	winLocalAppData := os.Getenv("LOCALAPPDATA") + "\\"

	switch installation.SoftwareName {
	case SoftwareVectorworks:
		return []string{
			winAppData + installation.SoftwareName + "\\" + installation.Year,
			winAppData + installation.SoftwareName + " " + installation.Year + " Installer",
			winAppData + installation.SoftwareName + " " + installation.Year + " Updater",
			winAppData + "Nemetschek\\" + installation.SoftwareName + "\\" + installation.Year,
			winAppData + "Nemetschek\\" + installation.SoftwareName + "\\accounts",
			winAppData + "Nemetschek\\" + installation.SoftwareName + " RMCache\\rm" + installation.Year,
			winAppData + "Nemetschek\\" + installation.SoftwareName + " Web Cache",
			winAppData + "vectorworks-installer",
			winAppData + "vectorworks-updater",
			winAppData + "vectorworks-updater-updater",
			winLocalAppData + "vectorworks-updater-updater",
			winLocalAppData + "Nemetschek",
		}
	case SoftwareVision:
		return []string{
			filepath.Join(winAppData, installation.SoftwareName, installation.Year),
			filepath.Join(winLocalAppData, "VisionUpdater"),
		}
	case SoftwareCloudServices:
		return []string{
			winAppData + "vectorworks-cloud-services-beta",
			winAppData + "vectorworks-cloud-services",
			winLocalAppData + "vectorworks-cloud-services-beta-updater",
		}
	}

	return nil
}

func (i Installation) Clean() {
	fmt.Println("Hello")
	for _, property := range i.Properties {
		k, _ := registry.OpenKey(registry.CURRENT_USER, property, registry.ALL_ACCESS)

		defer k.Close()

		names, _ := k.ReadSubKeyNames(-1)

		for _, name := range names {
			_ = registry.DeleteKey(k, name)
		}
		_ = registry.DeleteKey(k, "")
	}
	// TODO: Check for directory after as a way to verify deletion.

	for _, directory := range i.Directories {
		_ = os.RemoveAll(directory)
		// TODO: Implement error checking.
	}
}
