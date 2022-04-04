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
	fmt.Println("Adding Database scripts at : ", clonePath)
	err := commands.Clone(gitUrl, clonePath)
	if err != nil {
		return err
	}
	// run a go mod tidy
	fmt.Println("Tidying up the go mod file")
	err = commands.GoModTidy(currentDir)
	if err != nil {
		return err
	}
	return nil
}
