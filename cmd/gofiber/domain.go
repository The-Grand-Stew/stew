package gofiber

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"stew/pkg/commands"
	"stew/pkg/templates"
	"strings"
)

func AddModel(domains []string) error {
	currentDir, _ := os.Getwd()
	for _, modelName := range domains {
		addStruct(modelName)
		addControllers(modelName)
		addQueries(modelName)
		addRoutes(modelName)
		addMain(modelName)
	}

	err := commands.GoModTidy(currentDir)
	if err != nil {
		return err
	}
	err = commands.GoImports(currentDir)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func addStruct(modelName string) error {
	// Get path to add the model
	currentDir, _ := os.Getwd()
	directoryPath := filepath.Join(currentDir, "app", "models")
	// parse
	err := templates.AddGoFiberModelTemplate(modelName, directoryPath, templates.GoFiberModelTemplate)
	if err != nil {
		return err
	}
	return nil
}

func addControllers(modelName string) error {
	// Get path to add the model
	currentDir, _ := os.Getwd()
	directoryPath := filepath.Join(currentDir, "app", "controllers")
	// parse

	err := templates.AddGoFiberModelTemplate(modelName, directoryPath, templates.GoFiberControllerTemplate)
	if err != nil {
		return err
	}

	return nil

}

func addQueries(modelName string) error {
	// Get path to add the model
	currentDir, _ := os.Getwd()
	directoryPath := filepath.Join(currentDir, "app", "queries")
	// parse
	err := templates.AddGoFiberModelTemplate(modelName, directoryPath, templates.GoFiberQueryTemplate)
	if err != nil {
		return err
	}

	return nil
}

func addRoutes(modelName string) error {
	// Get path to add the model
	currentDir, _ := os.Getwd()
	directoryPath := filepath.Join(currentDir, "app", "routes")
	// parse
	err := templates.AddGoFiberModelTemplate(modelName, directoryPath, templates.GoFiberRouteTemplate)
	if err != nil {
		return err
	}
	return nil

}

func addMain(modelName string) error {
	// Get path to add the model
	currentDir, _ := os.Getwd()
	path := filepath.Join(currentDir, "cmd", "main.go")
	t, err := template.New("modelTemplate").Parse(templates.GoFiberMainTemplate)
	if err != nil {
		return err
	}
	// create output path to write the template
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	// vomit output to model file
	t.Execute(f, strings.Title(modelName))
	f.Close()
	return nil

}
