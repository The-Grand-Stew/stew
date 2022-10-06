package javaspringboot

import (
	"os"
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"ToUpper": strings.ToUpper,
	"ToLower": strings.ToLower,
	"Title":   strings.Title,
}

type DomainTemplate struct {
	AppName      string
	TemplateName string
	DomainName   string
	FilePath     string
	Method       string
}

func fileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}
	return false
}

func AddSpringBootTemplate(d DomainTemplate, skipIfFileExists bool, dirPath string, fileName string) error {
	if skipIfFileExists && fileExists(d.FilePath) {
		return nil // do nothing
	}
	_ = os.MkdirAll(dirPath, os.ModePerm) // create directory
	_, _ = os.Create(d.FilePath)          // create file
	d.DomainName = strings.Title(d.DomainName)
	d.AppName = strings.ToLower(d.AppName)
	t, err := template.New("modelTemplate").Funcs(funcMap).Parse(d.TemplateName)
	if err != nil {
		return err
	}
	// create output path to write the template
	f, err := os.Create(d.FilePath)
	if err != nil {
		return err
	}
	// vomit output to model file
	t.Execute(f, d)
	f.Close()
	return nil
}
