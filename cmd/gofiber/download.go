package gofiber

import (
	"fmt"
	"os"
	"path/filepath"
	"stew/pkg/commands"
	"stew/pkg/templates/repositories"
)

func DownloadTemplate(appName string) error {
	gitUrl := repositories.MicroservicesTemplates["go-fiber"]

	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	// clone template to path
	clonePath := filepath.Join(currentDir, appName)
	commands.ShowMessage("info", fmt.Sprintf("Cloning Template for go-fiber at location : %s", clonePath), true, true)
	err = commands.Clone(gitUrl, clonePath)
	if err != nil {
		return err
	}
	// do go mod init
	commands.ShowMessage("info", "Initialising a go mod init", true, true)
	err = commands.GoModInit(clonePath, appName)
	if err != nil {
		return err
	}
	// do a go mod tidy
	commands.ShowMessage("info", "Tidying up your go mod file", true, true)
	err = commands.GoModTidy(clonePath)
	if err != nil {
		return err
	}
	return nil
}
