package nodeexpress

import (
	"fmt"
	"os"
	"path/filepath"
	templates "stew/pkg/templates/nodeexpress"
	"strings"
)

var domainSettings = templates.DomainTemplate{AppName: "", DirectoryPath: ""}

func AddModel(appName string, domains []string) error {
	domainSettings.AppName = appName
	for _, modelName := range domains {
		addControllers(modelName)
		addRoutes(modelName)
		addSchema(modelName)
	}

	var httpMethods = []string{"post", "put", "get", "delete"}
	for _, method := range httpMethods {
		err = addModelFile(domain+"."+method, "test")
		if err != nil {
			fmt.Printf("Errored out %s", err)

			return err
		}
	}

	return nil
}

func addControllers(modelName string) error {
	// Get path to add the model
	currentDir, _ := os.Getwd()
	domainSettings.DirectoryPath = filepath.Join(currentDir, fileType, modelName+"s")
	err := os.MkdirAll(domainSettings.DirectoryPath, os.ModePerm)
	templateName := templates.NodeExpressRouteTemplate
	var method string
	switch fileType {
	case "route":
		templateName = templates.NodeExpressRouteTemplate
	case "controller":
		templateName = templates.NodeExpressControllerTemplate
	case "schema":
		templateName = templates.NodeExpressModelTemplate
	case "test":
		templateName = templates.NodeExpressTestTemplate
		splitModelName := strings.Split(modelName, ".")
		modelName = splitModelName[0]
		method = splitModelName[1]

	}

	return nil

}

func addSchema(modelName string) error {
	// Get path to add the model
	currentDir, _ := os.Getwd()
	domainSettings.DirectoryPath = filepath.Join(currentDir, "app", "schemas", modelName+"s")
	domainSettings.TemplateName = templates.NodeExpressModelTemplate
	domainSettings.DomainName = modelName
	//parse
	err := templates.AddNodeExpressapiTemplate(domainSettings)
	if err != nil {
		return err
	}

	return nil
}

func addRoutes(modelName string) error {
	// Get path to add the model
	currentDir, _ := os.Getwd()
	domainSettings.DirectoryPath = filepath.Join(currentDir, "app", "routes", modelName+"s")
	domainSettings.TemplateName = templates.NodeExpressRouteTemplate
	domainSettings.DomainName = modelName
	domainSettings.Method = method
	// parse
	err := templates.AddNodeExpressapiTemplate(domainSettings)
	if err != nil {
		return err
	}
	return nil

}
