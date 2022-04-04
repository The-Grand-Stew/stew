package gofiber

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"stew/pkg/commands"
	templates "stew/pkg/templates/gofiber"
	"strings"
)

var domainSettings = templates.DomainTemplate{AppName: "", DirectoryPath: ""}

func AddModel(appName string, domains []string) error {
	domainSettings.AppName = appName
	currentDir, _ := os.Getwd()
	for _, modelName := range domains {
		addStruct(modelName)
		addControllers(modelName)
		addQueries(modelName)
		addRoutes(modelName)
	}
	err := addMain(appName, domains)
	if err != nil {
		return err
	}
	err = commands.GoModTidy(currentDir)
	if err != nil {
		return err
	}
	// err = commands.GoImports(currentDir)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	return nil
}

func addStruct(modelName string) error {
	// Get path to add the model
	currentDir, _ := os.Getwd()
	domainSettings.DirectoryPath = filepath.Join(currentDir, "app", "models")
	domainSettings.TemplateName = templates.GoFiberModelTemplate
	domainSettings.DomainName = modelName
	// parse
	err := templates.AddGoFiberDomainTemplate(domainSettings)
	if err != nil {
		return err
	}
	return nil
}

func addControllers(modelName string) error {
	// Get path to add the model
	currentDir, _ := os.Getwd()
	// parse
	domainSettings.DirectoryPath = filepath.Join(currentDir, "app", "controllers")
	domainSettings.TemplateName = templates.GoFiberControllerTemplate
	domainSettings.DomainName = modelName

	err := templates.AddGoFiberDomainTemplate(domainSettings)
	if err != nil {
		return err
	}

	return nil

}

func addQueries(modelName string) error {
	// Get path to add the model
	currentDir, _ := os.Getwd()
	domainSettings.DirectoryPath = filepath.Join(currentDir, "app", "queries")
	domainSettings.TemplateName = templates.GoFiberQueryTemplate
	domainSettings.DomainName = modelName
	//parse
	err := templates.AddGoFiberDomainTemplate(domainSettings)
	if err != nil {
		return err
	}

	return nil
}

func addRoutes(modelName string) error {
	// Get path to add the model
	currentDir, _ := os.Getwd()
	domainSettings.DirectoryPath = filepath.Join(currentDir, "app", "routes")
	domainSettings.TemplateName = templates.GoFiberRouteTemplate
	domainSettings.DomainName = modelName
	// parse
	err := templates.AddGoFiberDomainTemplate(domainSettings)
	if err != nil {
		return err
	}
	return nil

}

func addMain(appName string, domains []string) error {
	// type routeTemplate struct {
	// 	appName string
	// 	routes  string
	// }

	// Get path to add the model
	var routes []string
	for _, domain := range domains {
		routes = append(routes, fmt.Sprintf(`routes.%sRoutes(app)`, strings.Title(domain)))
	}
	routeTemplate := map[string]interface{}{"appName": strings.ToLower(appName), "routes": strings.Join(routes, "\n")}
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
	t.Execute(f, routeTemplate)
	f.Close()
	return nil

}
