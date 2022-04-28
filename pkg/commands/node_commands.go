package commands

import (
	"os"
)

func ExecCommandWrapper(command string,options []string,directoryPath string) error {
	currentDir, _ := os.Getwd()
	os.Chdir(directoryPath)
	err := ExecCommand(command, options, true)
	if err != nil {
		return err
	}
	os.Chdir(currentDir)
	return nil
}

func NodeInit(directoryPath string) error {
	options := []string{"install"}
	command:="npm"
	return ExecCommandWrapper(command,options,directoryPath)
}

func NodeFormat(directoryPath string) error {
	options := []string{"run prettify"}
	command :="npm"
	return ExecCommandWrapper(command,options,directoryPath)
}


