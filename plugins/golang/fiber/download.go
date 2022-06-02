package fiber

import (
	"fmt"
	"os"
	"path/filepath"
	"stew/pkg/commands"
	"stew/pkg/logging"
	cmd "stew/plugins/golang/commands"
	"stew/plugins/golang/surveys"
)

func CreateMicroservice(appName, appPort string) error {
	currentDir, _ := os.Getwd()
	appPath := filepath.Join(currentDir, appName)

	gitUrl := surveys.FrameworkRepositories["fiber"]
	// clone template to path
	logging.ShowMessage("info", fmt.Sprintf("Cloning Template for  at location : %s", appPath), true, false)
	err := commands.Clone(gitUrl, appPath)
	if err != nil {
		return err
	}
	// do go mod init
	logging.ShowMessage("info", "Initialising a go mod init...", true, false)
	err = cmd.GoModInit(appPath, appName)
	if err != nil {
		return err
	}
	// do a go mod tidy
	logging.ShowMessage("info", "Tidying up your go mod file...", true, false)
	err = cmd.GoModTidy(appPath)
	if err != nil {
		return err
	}
	os.Chdir(appPath)
	err = AddModel(appName, appName, appPort)
	logging.ShowMessage("success", fmt.Sprintf("Go service %s created at %s and configured to run on port %s !", appName, appPath, appPort), true, true)
	os.Chdir(currentDir)
	return nil
}
