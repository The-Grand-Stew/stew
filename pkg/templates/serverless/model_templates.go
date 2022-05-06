package serverless

import (
	"html/template"
	"os"
	"path/filepath"
	"stew/pkg/utils"
	"strings"
)

var funcMap = template.FuncMap{
	"ToUpper": strings.ToUpper,
	"ToLower": strings.ToLower,
	"Title":   strings.Title,
}

type LambdaTemplate struct {
	AppName       string
	FunctionName  string
	DirectoryPath string
	Lang          string
	TemplateName  string
	PathPart      string
	HttpMethod    string
	HandlerName   string
	Environment   string
	Project       string
	Runtime       string
	Region        string
}

func AddServerlessTemplate(d LambdaTemplate, fileType string) error {
	d.FunctionName = strings.Title(d.FunctionName)
	var extension string
	filename := fileType
	secondLevelPath := "handlers"
	extension = utils.ExtensionMap[d.Lang]
	if fileType == "test" {
		filename = d.FunctionName + "_test"
		secondLevelPath = "__tests__"
	}
	path := filepath.Join(d.DirectoryPath, secondLevelPath, strings.ToLower(d.FunctionName))
	t, err := template.New("modelTemplate").Funcs(funcMap).Parse(d.TemplateName)
	if err != nil {
		return err
	}
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	// create output path to write the template
	fileOutpath := filepath.Join(path, filename+extension)
	f, err := os.Create(fileOutpath)
	if err != nil {
		return err
	}
	// vomit output to model file
	t.Execute(f, d)
	f.Close()
	return nil
}
