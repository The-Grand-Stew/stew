package commands

import (
	"os"
)

func GoModTidy(directoryPath string) error {
	currentDir, _ := os.Getwd()
	os.Chdir(directoryPath)
	options := []string{"mod", "tidy"}
	err := ExecCommand("go", options, true)
	if err != nil {
		return err
	}
	os.Chdir(currentDir)
	return nil

}

func GoModInit(directoryPath string, appName string) error {
	currentDir, _ := os.Getwd()
	os.Chdir(directoryPath)
	options := []string{"mod", "init", appName}
	err := ExecCommand("go", options, true)
	if err != nil {
		return err
	}
	os.Chdir(currentDir)
	return nil
}

func GoImports(directoryPath string) error {
	os.Chdir(directoryPath)
	options := []string{"-w", directoryPath}
	err := ExecCommand("goimports", options, true)
	if err != nil {
		return err
	}
	return nil
}

func GoFmt(directoryPath string) error {
	os.Chdir(directoryPath)
	options := []string{directoryPath}
	err := ExecCommand("gofmt", options, true)
	if err != nil {
		return err
	}
	return nil
}
