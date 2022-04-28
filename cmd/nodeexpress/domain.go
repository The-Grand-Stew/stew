package nodeexpress

import (
	"os"
	"path/filepath"
	"fmt"
	templates "stew/pkg/templates/nodeexpress"
)

var domainSettings = templates.DomainTemplate{AppName: "", DirectoryPath: ""}

func AddModel(appName string, domain string) error {
	domainSettings.AppName = appName
	err := addModelFile(domain,"route")
	if err != nil {
		fmt.Printf("Errored out %s",err)

		return err
	}
	err = addModelFile(domain,"controller")
	if err != nil {
		fmt.Printf("Errored out %s",err)

		return err
	}
	err = addModelFile(domain,"schema")
	if err != nil {
		fmt.Printf("Errored out %s",err)

		return err
	}
	return nil
}

func addModelFile(modelName string,fileType string) error{
	currentDir, _ := os.Getwd()
	domainSettings.DirectoryPath = filepath.Join(currentDir,fileType,modelName+"s")
	err := os.MkdirAll(domainSettings.DirectoryPath, os.ModePerm)
	templateName := templates.NodeExpressRouteTemplate
	switch fileType{
	case "route":
		templateName = templates.NodeExpressRouteTemplate
	case "controller":
		templateName = templates.NodeExpressControllerTemplate
	case "schema":
		templateName = templates.NodeExpressModelTemplate
	}
	domainSettings.TemplateName =templateName
	domainSettings.DomainName = modelName
	// parse
	err = templates.AddNodeExpressapiTemplate(domainSettings)
	if err != nil {
		return err
	}
	return nil
}
