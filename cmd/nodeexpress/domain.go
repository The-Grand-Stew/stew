package nodeexpress

import (
	"os"
	"path/filepath"
	templates "stew/pkg/templates/nodeexpress"
)

var domainSettings = templates.DomainTemplate{AppName: "", DirectoryPath: ""}

func AddModel(appName string, domains []string) error {
	domainSettings.AppName = appName
	for _, modelName := range domains {
		addControllers(modelName)
		addRoutes(modelName)
		addSchema(modelName)
	}
	return nil
}


func addControllers(modelName string) error {
	// Get path to add the model
	currentDir, _ := os.Getwd()
	// parse
	domainSettings.DirectoryPath = filepath.Join(currentDir, "app", "controllers",modelName+"s")
	domainSettings.TemplateName = templates.NodeExpressControllerTemplate
	domainSettings.DomainName = modelName

	err := templates.AddNodeExpressapiTemplate(domainSettings)
	if err != nil {
		return err
	}

	return nil

}

func addSchema(modelName string) error {
	// Get path to add the model
	currentDir, _ := os.Getwd()
	domainSettings.DirectoryPath = filepath.Join(currentDir, "app","schemas",modelName+"s")
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
	domainSettings.DirectoryPath = filepath.Join(currentDir, "app", "routes",modelName+"s")
	domainSettings.TemplateName = templates.NodeExpressRouteTemplate
	domainSettings.DomainName = modelName
	// parse
	err := templates.AddNodeExpressapiTemplate(domainSettings)
	if err != nil {
		return err
	}
	return nil

}


