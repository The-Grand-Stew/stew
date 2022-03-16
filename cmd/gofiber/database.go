package gofiber

import (
	"os"
	"path/filepath"
	exec "stew/pkg/commands"
	"stew/pkg/git"
)

func AddPostgres(appName string) error {
	currentDir, _ := os.Getwd()
	// Set project folder.
	clonePath := filepath.Join(currentDir, appName)
	// add gist to create db connection
	path := filepath.Join(clonePath, "platform", "database")
	git.Clone("go-postgres", path)
	// pick up connection credentials from env variables
	os.Chdir(appName)
	options := []string{"mod", "tidy"}
	err := exec.ExecCommand("go", options, true)
	if err != nil {
		exec.ShowError(err.Error())
	}
	os.Chdir(currentDir)
	return nil
}

func AddMySQL() error {
	// create directory for platform/db
	// add gist to create db connection
	// pick up connection credentials from env variables

	return nil
}
