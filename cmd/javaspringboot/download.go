package javaspringboot

import (
	"fmt"
	"os"
	"path/filepath"
	"stew/pkg/commands"
	"stew/pkg/templates/repositories"
)

func CreateMicroservice(appName string, appPort string) error {
	gitUrl := repositories.MicroservicesTemplates["java-springboot"]

	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	// clone template to path
	clonePath := filepath.Join(currentDir, appName)
	commands.ShowMessage("info", fmt.Sprintf("Cloning Template for java-springboot at location : %s", clonePath), true, false)

	err = commands.Clone(gitUrl, clonePath)
	if err != nil {
		return err
	}

	// in case we decide to build the code locally instead of doing that within the docker container
	// commands.ShowMessage("info", "Initializing the java project", true, false)
	// err = commands.JavaSpringBootInit(clonePath)
	// if err != nil {
	// 	return err
	// }
	// err = commands.SetAppPort(clonePath, appPort)
	// if err != nil {
	// 	return err
	// }

	os.Chdir(clonePath)
	err = AddModel(appName, appName)
	if err != nil {
		return err
	}
	os.Chdir(currentDir)
	commands.ShowMessage("success", fmt.Sprintf("Java service %s created at %s and configured to run on port %s !", appName, clonePath, appPort), true, true)
	return nil
}
