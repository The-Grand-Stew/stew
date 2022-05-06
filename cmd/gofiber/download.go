package gofiber

import (
	"fmt"
	"os"
	"path/filepath"
	"stew/pkg/commands"
)

func CreateMicroservice(appName, appPort string) error {
	currentDir, _ := os.Getwd()
	appPath := filepath.Join(currentDir, appName)
	err := commands.DownloadTemplate("go-fiber", appPath)
	if err != nil {
		return err
	}
	// do go mod init
	commands.ShowMessage("info", "Initialising a go mod init...", true, false)
	err = commands.GoModInit(appPath, appName)
	if err != nil {
		return err
	}
	// do a go mod tidy
	commands.ShowMessage("info", "Tidying up your go mod file...", true, false)
	err = commands.GoModTidy(appPath)
	if err != nil {
		return err
	}
	os.Chdir(appPath)
	err = AddModel(appName, appName, appPort)
	commands.ShowMessage("success", fmt.Sprintf("Go service %s created at %s and configured to run on port %s !", appName, appPath, appPort), true, true)
	os.Chdir(currentDir)
	return nil
}
