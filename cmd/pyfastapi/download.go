package pyfastapi

import (
	"fmt"
	"os"
	"path/filepath"
	"stew/pkg/commands"
	"stew/pkg/templates"
)

func DownloadTemplate(appName string) error {
	gitUrl := templates.MicroservicesTemplates["python-fastapi"]
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	// clone template to path
	clonePath := filepath.Join(currentDir, appName)
	fmt.Println("Cloning Template for python - fastapi at location : ", clonePath)
	err = commands.Clone(gitUrl, clonePath)
	if err != nil {
		return err
	}
	return nil
}
