package commands

import (
	"os"
)

func NodeInit(directoryPath string) error {
	currentDir, _ := os.Getwd()
	os.Chdir(directoryPath)
	options := []string{}
	err := ExecCommand("npm install", options, true)
	if err != nil {
		return err
	}
	os.Chdir(currentDir)
	return nil

}

func NodeFormat(directoryPath string) error {
	currentDir, _ := os.Getwd()
	os.Chdir(directoryPath)
	options := []string{}
	err := ExecCommand("npm run prettify", options, true)
	if err != nil {
		return err
	}
	os.Chdir(currentDir)
	return nil
}


