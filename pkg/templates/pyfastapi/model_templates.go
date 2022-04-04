package pyfastapi

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

func AddPythonFastapiTemplate(modelName string, directoryPath string, modelTemplate string) error {
	path := filepath.Join(directoryPath, strings.ToLower(modelName)+".py")
	t, err := template.New("modelTemplate").Funcs(funcMap).Parse(modelTemplate)
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
