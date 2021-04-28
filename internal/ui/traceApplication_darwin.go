package ui

import (
	"bufio"
	"bytes"
	"fmt"
	"howett.net/plist"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
)

type plistOptions struct {
	properties map[string]interface{}
}

func runApplication(ch chan []byte, targetFile string) {
	var err error

	targetFile, err = getContentTargetFile(targetFile)
	if err != nil {
		err = fmt.Errorf("runApplicationError: %v", err)
		panic(err)
	}
	cmd := exec.Command(targetFile)

	outReader, outWriter := io.Pipe()
	errReader, errWriter := io.Pipe()

	cmd.Stdout = outWriter
	cmd.Stderr = errWriter

	//mw := io.MultiWriter(outWriter, errWriter, &buffer)
	mr := io.MultiReader(outReader, errReader, &buffer)

	// Is the transfer the output of the application, through the channel.

	go func() {
		reader := bufio.NewReader(mr)

		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				break
			}
			ch <- line
			ch <- []byte{'\n'}
		}
	}()

	if err := cmd.Run(); err != nil {
		log.Panicln(err)
	}
}

func getContentTargetFile(targetPath string) (string, error) {
	var plistData plistOptions

	// Create Info.plist target string and read in the file
	plistFileString := filepath.Join(targetPath, "Contents", "Info.plist")
	plistFile, err := ioutil.ReadFile(plistFileString)
	if err != nil {
		return "", fmt.Errorf("error reading in info.plist file for selected application - %v", err)
	}

	// read and decode plist file
	plistReader := bytes.NewReader(plistFile)
	decoder := plist.NewDecoder(plistReader)
	err = decoder.Decode(&plistData.properties)
	if err != nil {
		return "", fmt.Errorf("error decoding provided info.plist file - %v", err)
	}

	// get the target binary from the plist file and construct a real path to return
	targetBinaryString := plistData.properties["CFBundleExecutable"].(string)
	out := filepath.Join(targetPath, "Contents", "MacOS", targetBinaryString)

	return out, nil
}
