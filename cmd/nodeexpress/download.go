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
	commands.ShowMessage("info", fmt.Sprintf("Failed to clone repo : %s", err), true, true)
		return err
	}
	// do npm install
	commands.ShowMessage("info", "Initialising the nodejs project", true, true)
	err = commands.NodeInit(clonePath)
	if err != nil {
		commands.ShowMessage("error", fmt.Sprintf("Failed to initialize repo : %s", err), true, true)
		return err
	}
	// run prettifier
	// commands.ShowMessage("info", "Prettifying your code", true, true)
	// err = commands.NodeFormat(clonePath)
	// if err != nil {
	// 	commands.ShowMessage("error", fmt.Sprintf("Failed to format code in repo : %s", err), true, true)
	// 	return err
	// }
	return nil
}
