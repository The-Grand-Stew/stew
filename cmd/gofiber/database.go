package gofiber

import (
	"fmt"
	"os"
	"path/filepath"
	"stew/pkg/commands"
	"stew/pkg/templates"
)

func AddPostgres(appName string) error {
	// get template for postgres
	gitUrl := templates.MicroservicesTemplates["go-fiber-postgres"]
	currentDir, _ := os.Getwd()
	// clone gist to db folder
	clonePath := filepath.Join(currentDir, appName, "platform", "database")
	fmt.Println("Adding Database scripts at : ", clonePath)
	err := commands.Clone(gitUrl, clonePath)
	if err != nil {
		return err
	}
	// run a go mod tidy
	fmt.Println("Tidying up the go mod file")
	err = commands.GoModTidy(filepath.Join(currentDir, appName))
	if err != nil {
		return err
	}
	return nil
}
