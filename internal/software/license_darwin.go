package software

import (
	"bufio"
	"bytes"
	"fmt"
	"howett.net/plist"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type LicenseOpts struct {
	serial map[string]string `plist:"NNA User License"`
}

func GetSerialLocation(installation Installation) string {
	switch installation.SoftwareName {
	case SoftwareVectorworks:
		return HomeDirectory + "/Library/Preferences/net.nemetschek.vectorworks.license." + installation.Year + ".plist"
	case SoftwareVision:
		return HomeDirectory + "/Library/Preferences/net.vectorworks.vision.license." + installation.Year + ".plist"
	}

	return ""
}

// getSerial will read in a plist, decode it and return a keyed value as a string value
func getSerial(installation Installation) string {
	serialLocation := GetSerialLocation(installation)

	// Read in plist
	plistFile, err := ioutil.ReadFile(serialLocation)
	buffer := bytes.NewReader(plistFile)
	Check(err)

	// parse and return plist serial
	var plistData LicenseOpts
	decoder := plist.NewDecoder(buffer)
	err = decoder.Decode(&plistData.serial)
	Check(err)

	return plistData.serial[`NNA User License`]
}

// replaceOldSerial
func replaceOldSerial(installation Installation, newSerial string) {
	licenseLocation := GetSerialLocation(installation)
	plistFile, err := os.Open(licenseLocation)
	Check(err)
	err = plistFile.Truncate(0)

	newSerial = cleanSerial(newSerial) // Clean and verify serial

	plistData := &LicenseOpts{
		serial: map[string]string{"NNA User License": newSerial},
	}

	fmt.Println(plistData.serial)
	buffer := &bytes.Buffer{}
	encoder := plist.NewEncoder(buffer)

	err = encoder.Encode(plistData.serial)
	Check(err)

	err = os.WriteFile(licenseLocation, buffer.Bytes(), 0644)
	Check(err)

	w := bufio.NewWriter(buffer)
	n4, err := w.WriteString("buffered\n")
	Check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	err = w.Flush()
	Check(err)

	refreshPList()
}

func refreshPList() {
	fmt.Println("Refreshing plist files...")
	// osascript -e 'do shell script "sudo killall -u $USER cfprefsd" with administrator privileges'
	cmd := exec.Command(`osascript`, "-s", "h", "-e", `do shell script "sudo killall -u $USER cfprefsd" with administrator privileges`)
	stderr, err := cmd.StderrPipe()
	log.SetOutput(os.Stderr)
	Check(err)

	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}

	cmdErrOutput, _ := ioutil.ReadAll(stderr)
	fmt.Printf("%s\n", cmdErrOutput)

	if err = cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
