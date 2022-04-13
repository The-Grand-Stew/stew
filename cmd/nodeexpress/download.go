package nodeexpress

import (
	"fmt"
	"os"
	"path/filepath"
	"stew/pkg/commands"
	"stew/pkg/templates/repositories"
)

func DownloadTemplate(appName string) error {
	gitUrl := repositories.MicroservicesTemplates["node-express"]

	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	// clone template to path
	clonePath := filepath.Join(currentDir, appName)
	commands.ShowMessage("info", fmt.Sprintf("Cloning Template for node-express at location : %s", clonePath), true, true)
	err = commands.Clone(gitUrl, clonePath)
	if err != nil {
		return err
	}
	// do npm install
	commands.ShowMessage("info", "Running an npm install", true, true)
	err = commands.NpmInstall(clonePath)
	if err != nil {
		return err
	}
	return nil
}
