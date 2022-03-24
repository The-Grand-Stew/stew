package gofiber

import (
	"fmt"
	"os"
	"path/filepath"
	"stew/pkg/commands"
	"stew/pkg/templates"
)

func DownloadTemplate(appName string) error {
	gitUrl := templates.MicroservicesTemplates["go-fiber"]

	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	// clone template to path
	clonePath := filepath.Join(currentDir, appName)
	fmt.Println("Cloning Template for go-fiber at location : ", clonePath)
	err = commands.Clone(gitUrl, clonePath)
	if err != nil {
		return err
	}
	// do go mod init
	fmt.Println("Initialising a go mod init")
	err = commands.GoModInit(clonePath, appName)
	if err != nil {
		return err
	}
	// do a go mod tidy
	fmt.Println("Tidying up your go mod file")
	err = commands.GoModTidy(clonePath)
	if err != nil {
		return err
	}
	return nil
}
