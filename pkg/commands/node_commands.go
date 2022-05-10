package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ExecCommandWrapper(command string, options []string, directoryPath string) error {
	currentDir, _ := os.Getwd()
	os.Chdir(directoryPath)
	err := ExecCommand(command, options, true)
	if err != nil {
		os.Chdir(currentDir)
		return err
	}
	os.Chdir(currentDir)
	return nil
}

func NodeInit(directoryPath string) error {
	options := []string{"install"}
	command := "npm"
	return ExecCommandWrapper(command, options, directoryPath)
}

func NodeFormat(directoryPath string) error {
	options := []string{"run prettify"}
	command := "npm"
	return ExecCommandWrapper(command, options, directoryPath)
}

func SetAppPort(directoryPath string, appPort string) error {
	currentDir, _ := os.Getwd()
	clonePath := filepath.Join(directoryPath, "bin")
	os.Chdir(clonePath)
	input, err := ioutil.ReadFile("www")
	if err != nil {
		fmt.Printf("%s", err)
		return err
	}
	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		if strings.Contains(line, "normalizePort(process.env.PORT || \"3000\")") {
			lines[i] = "const port = normalizePort(process.env.PORT || " + "\"" + appPort + "\");"
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile("www", []byte(output), 0644)
	if err != nil {
		os.Chdir(currentDir)
		return err
	}

	os.Chdir(currentDir)
	return nil
}
