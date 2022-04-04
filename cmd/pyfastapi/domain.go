package pyfastapi

import (
	"os"
	"path/filepath"
	"stew/pkg/templates/pyfastapi"
)

//TODO: ADD CRUD FUNCTIONALITIES FOR SCHEMA IF DB IS PRESENT

func AddModel(domains []string) error {

	for _, modelName := range domains {
		addSchema(modelName)
	}
	return nil
}

func addSchema(modelName string) error {
	// Get path to add the model
	currentDir, _ := os.Getwd()
	directoryPath := filepath.Join(currentDir, "app", "schemas")
	// parse
	err := pyfastapi.AddPythonFastapiTemplate(modelName, directoryPath, pyfastapi.PyFastAPISchema)
	if err != nil {
		return err
	}
	return nil
}
