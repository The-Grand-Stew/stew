package nodeexpress

import (
	"fmt"
	"os"
	"path/filepath"
	"stew/pkg/commands"
	"stew/pkg/templates/repositories"
)

func CreateMicroservice(appName string, appPort string) error {
	gitUrl := repositories.MicroservicesTemplates["node-express"]

	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	// clone template to path
	clonePath := filepath.Join(currentDir, appName)
	commands.ShowMessage("info", fmt.Sprintf("Cloning Template for node-express at location : %s", clonePath), true, false)
	err = commands.Clone(gitUrl, clonePath)
	if err != nil {
		return err
	}
	// do npm install
	commands.ShowMessage("info", "Initializing the nodejs project", true, false)
	err = commands.NodeInit(clonePath)
	if err != nil {
		return err
	}
	err = commands.SetAppPort(clonePath, appPort)
	if err != nil {
		return err
	}
	os.Chdir(clonePath)
	err = AddModel(appName, appName)
	if err != nil {
		return err
	}
	os.Chdir(currentDir)
	commands.ShowMessage("success", fmt.Sprintf("Node service %s created at %s and configured to run on port %s !", appName, clonePath, appPort), true, true)
	return nil
}
