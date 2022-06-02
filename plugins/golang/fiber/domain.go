package fiber

import (
	"html/template"
	"os"
	"path/filepath"
	"stew/pkg/configs"
	cmd "stew/plugins/golang/commands"
)

var funcMap = template.FuncMap(configs.FuncMap)

func AddModel(domain, appName, appPort string) error {
	// get the source
	currentDir, _ := os.Getwd()
	// TODO: go routine-ify this thing. for loop not good
	for _, scaffold := range Scaffold {
		filename := scaffold.Filename
		if filename == "" {
			filename = domain + Extension
		}
		destinationPath := filepath.Join(currentDir, scaffold.Destination, filename)
		templateName := scaffold.Source
		t, err := template.New("modelTemplate").Funcs(funcMap).Parse(templateName)
		if err != nil {
			return err
		}
		// create output path to write the template
		f, err := os.Create(destinationPath)
		if err != nil {
			return err
		}
		domainSettings := map[string]string{"AppName": appName, "DomainName": domain, "AppPort": appPort}
		// vomit output to model file
		t.Execute(f, domainSettings)
		f.Close()
	}
	err := cmd.GoModTidy(currentDir)
	if err != nil {
		return err
	}
	return nil
}
