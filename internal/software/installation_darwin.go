package software

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func FindInstallationYears(softwareLabel SoftwareLabel) []string {
	var years []string

	// FIXME: new method to get plist, current limits the returned value
	files, err := ioutil.ReadDir(GetHomeDir() + "/Library/Preferences") // gets list of all plist file names
	Check(err)

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

	return years
}

func findProperties(installation Installation) []string {
	switch installation.SoftwareLabel {
	case SoftwareVectorworks:
		return []string{
			"net.nemetschek.vectorworks.license." + installation.Year + ".plist",
			"net.nemetschek.vectorworks." + installation.Year + ".plist",
			"net.nemetschek.vectorworks.spotlightimporter.plist",
			"net.nemetschek.vectorworks.plist",
			"net.nemetschek.vectorworksinstaller.helper.plist",
			"net.nemetschek.vectorworksinstaller.plist",
			"net.vectorworks.vectorworks." + installation.Year + ".plist",
		}
	case SoftwareVision:
		return []string{
			"com.qtproject.plist",
			"com.vwvision.Vision" + installation.Year + ".plist",
			"com.yourcompany.Vision.plist",
			"net.vectorworks.Vision.plist",
			"net.vectorworks.vision.license." + installation.Year + ".plist",
		}
	}

	return nil
}

func findDirectories(installation Installation) []string {
	switch installation.SoftwareLabel {
	case SoftwareVectorworks:
		return []string{
			GetHomeDir() + "/Library/Application\\ Support/Vectorworks\\ RMCache/rm" + installation.Year,
			GetHomeDir() + "/Library/Application\\ Support/Vectorworks\\ Cloud\\ Services",
			GetHomeDir() + "/Library/Application\\ Support/Vectorworks/" + installation.Year,
			GetHomeDir() + "/Library/Application\\ Support/vectorworks-installer-wrapper",
		}
	case SoftwareVision:
		return []string{
			GetHomeDir() + "/Library/Application\\ Support/Vision/" + installation.Year,
			GetHomeDir() + "/Library/Application\\ Support/VisionUpdater",
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

	return nil
}

func (i Installation) Clean() {
	plistPath := GetHomeDir() + "/Library/Preferences/"
	// Deletes relevant plist files for select software/version
	for _, plist := range i.Properties {
		_ = os.RemoveAll(plistPath + plist)
		//TODO: Add logging for user feedback.
	}
	// TODO: Check for directory after as a way to verify deletion.

	for _, directory := range i.Directories {
		_ = os.RemoveAll(directory)
		// TODO: Implement error checking.
	}
}
