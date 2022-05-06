package serverless

import (
	"fmt"
	"os"
	"path/filepath"
	"stew/pkg/commands"
	"stew/pkg/templates/repositories"
)

func AddPostgres() error {
	// get template for postgres
	gitUrl := repositories.MicroservicesTemplates["node-express-utils"]
	currentDir, _ := os.Getwd()
	// clone gist to db folder
	clonePath := filepath.Join(currentDir, "database")
	commands.ShowMessage("info", fmt.Sprintf("Adding Database scripts at : %s", clonePath), true, true)
	err := commands.Clone(gitUrl, clonePath)
	if err != nil {
		fmt.Printf("Failed to clone utils repo %s", err)
		return err
	}
	return nil
}
