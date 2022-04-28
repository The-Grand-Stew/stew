package gofiber

import (
	"os"
	"path/filepath"
	"stew/pkg/commands"
)

func CreateMicroservice(appName string) error {
	currentDir, _ := os.Getwd()
	appPath := filepath.Join(currentDir, appName)
	err := commands.DownloadTemplate("go-fiber", appPath)
	if err != nil {
		return err
	}
	// do go mod init
	commands.ShowMessage("info", "Initialising a go mod init", true, true)
	err = commands.GoModInit(appPath, appName)
	if err != nil {
		return err
	}
	// do a go mod tidy
	commands.ShowMessage("info", "Tidying up your go mod file", true, true)
	err = commands.GoModTidy(appPath)
	if err != nil {
		return err
	}
	os.Chdir(appPath)
	err = AddModel(appName, appName)
	os.Chdir(currentDir)
	return nil
}
