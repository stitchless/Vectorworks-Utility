package software

import (
	"errors"
	"os"
	"regexp"
	"strings"
)

// TODO: Replace all home directories with GetConfigDirectory

func FindInstallationYears(softwareLabel SoftwareName) ([]string, error) {
	var years []string

	files, err := os.ReadDir(HomeDirectory + "/Library/Preferences") // gets list of all plist file names
	if err != nil {
		return nil, errors.New("error: could not get files from local library/Preferences")
	}

	// returns all license year numbers found in plist file names from the files variable
	for _, f := range files {
		file := strings.Contains(f.Name(), strings.ToLower(softwareLabel+".license."))
		if file {
			year := regexp.MustCompile("[0-9]+").FindString(f.Name())
			if year != "" {
				years = append(years, year)
			}
		}
	}

	return years, nil
}

// setProperties will take in an installation and assign it's properties strings
func (installation *Installation) setProperties() {
	switch installation.SoftwareName {
	case SoftwareVectorworks:
		installation.Properties = []string{
			"net.nemetschek.vectorworks.license." + installation.Year + ".plist",
			"net.nemetschek.vectorworks." + installation.Year + ".plist",
			"net.nemetschek.vectorworks.spotlightimporter.plist",
			"net.nemetschek.vectorworks.plist",
			"net.nemetschek.vectorworksinstaller.helper.plist",
			"net.nemetschek.vectorworksinstaller.plist",
			"net.vectorworks.vectorworks." + installation.Year + ".plist",
		}
	case SoftwareVision:
		installation.Properties = []string{
			"com.qtproject.plist",
			"com.vwvision.Vision" + installation.Year + ".plist",
			"com.yourcompany.Vision.plist",
			"net.vectorworks.Vision.plist",
			"net.vectorworks.vision.license." + installation.Year + ".plist",
		}
	}
}

// setUserData well set all user data based on the target software
func (installation *Installation) setUserData() {
	switch installation.SoftwareName {
	case SoftwareVectorworks:
		installation.Directories = []string{
			HomeDirectory + "/Library/Application\\ Support/Vectorworks\\ RMCache/rm" + installation.Year,
			HomeDirectory + "/Library/Application\\ Support/Vectorworks\\ Cloud\\ Services",
			HomeDirectory + "/Library/Application\\ Support/Vectorworks/" + installation.Year,
			HomeDirectory + "/Library/Application\\ Support/vectorworks-installer-wrapper",
		}
	case SoftwareVision:
		installation.Directories = []string{
			HomeDirectory + "/Library/Application\\ Support/Vision/" + installation.Year,
			HomeDirectory + "/Library/Application\\ Support/VisionUpdater",
			"/Library/Frameworks/QtConcurrent.framework",
			"/Library/Frameworks/QtCore.framework",
			"/Library/Frameworks/QtDBus.framework",
			"/Library/Frameworks/QtGui.framework",
			"/Library/Frameworks/QtNetwork.framework",
			"/Library/Frameworks/QtOpenGL.framework",
			"/Library/Frameworks/QtPlugins",
			"/Library/Frameworks/QtPositioning.framework",
			"/Library/Frameworks/QtPrintSupport.framework",
			"/Library/Frameworks/QtQml.framework",
			"/Library/Frameworks/QtQuick.framework",
			"/Library/Frameworks/QtWebChannel.framework",
			"/Library/Frameworks/QtWebEngine.framework",
			"/Library/Frameworks/QtWebEngineCore.framework",
			"/Library/Frameworks/QtWebEngineWidgets.framework",
			"/Library/Frameworks/QtWidgets.framework",
			"/Library/Frameworks/QtXml.framework",
			"/Library/Frameworks/rpath_manipulator.sh",
			"/Library/Frameworks/setup_qt_frameworks.sh",
		}
	}
}

// setRMCache sets the system path for the resource manager cache directory
func (installation *Installation) setRMCache() {
	installation.RMCache = HomeDirectory + "/Library/Application\\ Support/Vectorworks\\ RMCache/rm" + installation.Year
}

// setLogFiles sets all the log files paths for the target software
func (installation *Installation) setLogFiles() {
	installation.LogFiles = []string{
		HomeDirectory + "/Library/Application\\ Support/Vectorworks/" + installation.Year + "/VW User Log Sent.txt",
		HomeDirectory + "/Library/Application\\ Support/Vectorworks/" + installation.Year + "/VW User Log.txt",
	}
}

func (installation Installation) Clean() {
	plistPath := HomeDirectory + "/Library/Preferences/"
	// Deletes relevant plist files for select software/version
	for _, plist := range installation.Properties {
		err := os.RemoveAll(plistPath + plist)
		if err != nil {
			errors.New("error: could not remove the plist file: " + plistPath + plist)
		}
	}

	for _, directory := range installation.Directories {
		err := os.RemoveAll(directory)
		if err != nil {
			errors.New("error: could not delete the directory: " + directory)
		}
	}
}