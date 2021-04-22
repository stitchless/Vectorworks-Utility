package software

import (
	"log"
	"os"
)

// Check Error Checking
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// GetHomeDirectory Define users home directory
func GetHomeDirectory() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return home
}

func GetConfigDirectory() (string, error) {
	config, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return config, nil
}