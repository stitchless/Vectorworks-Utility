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
	var plistFile []byte
	var plistData plistOptions


	// Create Info.plist target string
	plistFileString := filepath.Join(targetFile, "Contents", "Info.plist")
	// Read in plist
	plistFile, err = ioutil.ReadFile(plistFileString)
	if err != nil {
		fmt.Errorf("error reading plist file: %v", err)
	}


	plistReader := bytes.NewReader(plistFile)
	// parse and return plist serial
	decoder := plist.NewDecoder(plistReader)
	//
	err = decoder.Decode(&plistData.properties)

	targetFile = filepath.Join(targetFile, "Contents", "MacOS", plistData.properties["CFBundleExecutable"].(string))
	fmt.Println(targetFile)

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
