package gofiber

import (
	"fmt"
	"os"
	"path/filepath"
	"stew/pkg/commands"
	"stew/pkg/templates/repositories"
)

func AddPostgres() error {
	// get template for postgres
	gitUrl := repositories.MicroservicesTemplates["go-fiber-postgres"]
	currentDir, _ := os.Getwd()
	// clone gist to db folder
	clonePath := filepath.Join(currentDir, "platform", "database")
	commands.ShowMessage("info", fmt.Sprintf("Adding Database scripts at : %s", clonePath), true, true)
	err := commands.Clone(gitUrl, clonePath)
	if err != nil {
		return err
	}
	// run a go mod tidy
	commands.ShowMessage("info", "Tidying up the go mod file", true, true)
	err = commands.GoModTidy(currentDir)
	if err != nil {
		return err
	}
	return nil
}
