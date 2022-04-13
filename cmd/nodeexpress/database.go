package nodeexpress

import (
	"os"
	"stew/pkg/commands"
)

func AddDatabase(appName string, databaseName string) error {
	currentDir, _ := os.Getwd()
	err = commands.NpmRunUtils(currentDir, databaseName)
	if err != nil {
		return err
	}
	return nil
}
