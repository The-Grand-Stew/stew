package gofiber

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
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
}

func AddGoFiberDomainTemplate(d DomainTemplate) error {
	d.DomainName = strings.Title(d.DomainName)
	d.AppName = strings.ToLower(d.AppName)
	path := filepath.Join(d.DirectoryPath, strings.ToLower(d.DomainName)+".go")
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
