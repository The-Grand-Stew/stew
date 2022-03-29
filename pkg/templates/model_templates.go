package templates

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"ToUpper": strings.ToUpper,
	"ToLower": strings.ToLower,
}

func AddGoFiberModelTemplate(modelName string, directoryPath string, modelTemplate string) error {
	path := filepath.Join(directoryPath, strings.ToLower(modelName)+".go")
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
