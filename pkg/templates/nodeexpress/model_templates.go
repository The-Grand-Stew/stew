package nodeexpress

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

var funcMap = template.FuncMap{
	"ToUpper": strings.ToUpper,
	"ToLower": strings.ToLower,
	"Title":   strings.Title,
}

type DomainTemplate struct {
	AppName       string
	TemplateName  string
	DomainName    string
	DirectoryPath string
	Method        string
}

// func AddNodeExpressapiTemplate(modelName string, directoryPath string, modelTemplate string) error {
// 	path := filepath.Join(directoryPath, strings.ToLower(modelName)+".js")
// 	t, err := template.New("modelTemplate").Funcs(funcMap).Parse(modelTemplate)
// 	if err != nil {
// 		return err
// 	}
// 	// create output path to write the template
// 	f, err := os.Create(path)
// 	if err != nil {
// 		return err
// 	}
// 	// vomit output to model file
// 	t.Execute(f, strings.Title(modelName))
// 	f.Close()
// 	return nil
// }

func AddNodeExpressapiTemplate(d DomainTemplate) error {
	d.DomainName = strings.Title(d.DomainName)
	d.AppName = strings.ToLower(d.AppName)
	path := filepath.Join(d.DirectoryPath, strings.ToLower(d.DomainName)+".js")
	t, err := template.New("modelTemplate").Funcs(funcMap).Parse(d.TemplateName)
	if err != nil {
		return err
	}
	// create output path to write the template
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	// vomit output to model file
	t.Execute(f, d)
	f.Close()
	return nil
}
