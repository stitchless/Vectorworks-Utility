package software

import (
	"golang.org/x/sys/windows/registry"
	"log"
	"strconv"
	"strings"
)

// getSerial will search the registry for any valid serials.
func getSerial(installation Installation) string {
	serialLocation := getSerialLocation(installation)

	// Get the Registry Key
	key, _ := registry.OpenKey(registry.CURRENT_USER, serialLocation, registry.QUERY_VALUE)
	defer func() {
		_ = key.Close()
	}()

	switch installation.Software {
	case SoftwareVectorworks:
		serial, _, _ := key.GetStringValue("User Serial Number")
		return serial
	case SoftwareVision:
		serial, _, _ := key.GetStringValue("")
		return serial
	}

	return ""
}

// convertYearToVersion returns a version number as opposed to a version year.
// This is need in the case of a registry due to application versions being used
// in place of version years
func convertYearToVersion(appYear string) string {
	values := strings.Split(appYear, "")[2:4]
	appVersion := values[0] + values[1]
	i, err := strconv.Atoi(appVersion)
	if err != nil {
		log.Fatal(err)
	}
	versionMath := i + 5
	appVersion = strconv.Itoa(versionMath)
	return appVersion
}

func getSerialLocation(installation Installation) string {
	version := convertYearToVersion(installation.Year)

	switch installation.Software {
	case SoftwareVectorworks:
		return "SOFTWARE\\Nemetschek\\Vectorworks " + version + "\\Registration"
	case SoftwareVision:
		return "SOFTWARE\\VectorWorks\\Vision " + installation.Year + "\\Registration"
	}

	return ""
}

func replaceOldSerial(installation Installation, newSerial string) {
	serialLocation := getSerialLocation(installation)

	key, err := registry.OpenKey(registry.CURRENT_USER, serialLocation, registry.SET_VALUE)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = key.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = key.SetStringValue("User Serial Number", newSerial)
	if err != nil {
		log.Fatal(err)
	}
}

//func cleanSerial(serial string) string {
//	// TODO: Clear empty space (Done)
//	// TODO: REGEX confirmation
//	// TODO: Return error and cancel replacement
//	r := regexp.MustCompile(`(.{6})-(.{6})-(.{6})-(.{6})`)
//	serial = strings.TrimSpace(serial)
//	parseSerial := r.FindAllString(serial, -1)
//	return parseSerial[0]
//}
