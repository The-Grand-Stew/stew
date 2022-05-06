package nodeexpress

import (
	"os"
	"path/filepath"
	templates "stew/pkg/templates/nodeexpress"
	"strings"
)

var domainSettings = templates.DomainTemplate{AppName: "", DirectoryPath: ""}

func AddModel(appName string, domain string) error {
	domainSettings.AppName = appName
	err := addModelFile(domain, "route")
	if err != nil {

		return err
	}
	err = addModelFile(domain, "controller")
	if err != nil {

		return err
	}
	err = addModelFile(domain, "schema")
	if err != nil {

		return err
	}

	var httpMethods = []string{"post", "put", "get", "delete"}
	for _, method := range httpMethods {
		err = addModelFile(domain+"."+method, "test")
		if err != nil {

			return err
		}
	}

	return nil
}

func addModelFile(modelName string, fileType string) error {
	currentDir, _ := os.Getwd()
	domainSettings.DirectoryPath = filepath.Join(currentDir, fileType, modelName)
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
	domainSettings.TemplateName = templateName
	domainSettings.DomainName = modelName
	domainSettings.Method = method
	// parse
	err = templates.AddNodeExpressapiTemplate(domainSettings)
	if err != nil {
		return err
	}
	return nil
}
