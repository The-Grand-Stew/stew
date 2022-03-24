package gofiber

import (
	"os"
	"path/filepath"
	"stew/pkg/commands"
	"stew/pkg/templates"
)

const queryTemplate string = `package queries`

const controllerTemplate string = `package controllers`

const routeTemplate string = `package routes`

const mainTemplate string = `package main`

func AddModel(modelName, appName string) error {
	addStruct(modelName, appName)
	addControllers(modelName, appName)
	addQueries(modelName, appName)
	addRoutes(modelName, appName)
	return nil
}

func addStruct(modelName string, appName string) error {
	// Get path to add the model
	currentDir, _ := os.Getwd()
	directoryPath := filepath.Join(currentDir, appName, "app", "models")
	// parse
	err := templates.AddGoFiberModelTemplate(modelName, directoryPath, templates.GoFiberModelTemplate)
	if err != nil {
		return err
	}
	return nil
}

func addControllers(modelName, appName string) error {
	// Get path to add the model
	currentDir, _ := os.Getwd()
	directoryPath := filepath.Join(currentDir, appName, "app", "controllers")
	// parse
	err := templates.AddGoFiberModelTemplate(modelName, directoryPath, templates.GoFiberControllerTemplate)
	if err != nil {
		return err
	}
	err = commands.GoModTidy(filepath.Join(currentDir, appName))
	if err != nil {
		return err
	}
	return nil

}

func addQueries(modelName, appName string) error {
	// Get path to add the model
	currentDir, _ := os.Getwd()
	directoryPath := filepath.Join(currentDir, appName, "app", "queries")
	// parse
	err := templates.AddGoFiberModelTemplate(modelName, directoryPath, templates.GoFiberQueryTemplate)
	if err != nil {
		return err
	}
	err = commands.GoModTidy(filepath.Join(currentDir, appName))
	if err != nil {
		return err
	}
	return nil
}

func addRoutes(modelName, appName string) error {
	// Get path to add the model
	currentDir, _ := os.Getwd()
	directoryPath := filepath.Join(currentDir, appName, "app", "routes")
	// parse
	err := templates.AddGoFiberModelTemplate(modelName, directoryPath, templates.GoFiberRouteTemplate)
	if err != nil {
		return err
	}
	err = commands.GoModTidy(filepath.Join(currentDir, appName))
	if err != nil {
		return err
	}
	return nil

}

func addMain(modelName, appName string) error {
	return nil

}
