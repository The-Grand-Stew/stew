package nodeexpress

import (
	"os"
	"stew/pkg/commands"
)

var err error

func AddModel(appName string, domains []string) error {
	currentDir, _ := os.Getwd()
	for _, modelName := range domains {
		commands.NpmRunModel(currentDir, modelName)
	}
	if err != nil {
		return err
	}
	return nil
}
