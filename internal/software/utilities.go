package software

import (
	"log"
	"os"
)

// Error Checking
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Define users home directory
func GetWD() string {
	var err error
	var dir string
	dir, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

// Define users home directory
func GetHomeDir() string {
	var err error
	var dir string
	dir, err = os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}