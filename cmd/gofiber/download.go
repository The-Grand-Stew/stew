package gofiber

import (
	"os"
	"path/filepath"
	exec "stew/pkg/commands"
	"stew/pkg/git"
)

func DownloadTemplate(appName string) error {
	// get git url
	// clone to specific folder = app name
	currentDir, _ := os.Getwd()
	// Set project folder.
	clonePath := filepath.Join(currentDir, appName)
	git.Clone("go-fiber", clonePath)
	os.Chdir(appName)
	// run a go mod init inside with app name
	options := []string{"mod", "init", appName}
	err := exec.ExecCommand("go", options, true)
	if err != nil {
		exec.ShowError(err.Error())
	}
	// run a go mod tidy
	options = []string{"mod", "tidy"}
	err = exec.ExecCommand("go", options, true)
	if err != nil {
		exec.ShowError(err.Error())
	}
	os.Chdir(currentDir)
	return nil
}
