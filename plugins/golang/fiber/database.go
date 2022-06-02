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

func AddDatabase(dbName string) error {
	currentDir, _ := os.Getwd()
	gitUrl := surveys.DatabaseRepositories[dbName]
	clonePath := filepath.Join(currentDir, "platform", "database")
	// clone template to path
	logging.ShowMessage("info", fmt.Sprintf("Adding Database scripts at : %s", clonePath), true, false)
	err := commands.Clone(gitUrl, clonePath)
	if err != nil {
		return err
	}
	// run a go mod tidy
	logging.ShowMessage("info", "Tidying up the go mod file...", true, false)
	err = cmd.GoModTidy(currentDir)
	if err != nil {
		return err
	}
	return nil
}
