package javaspringboot

// TODO: so far I've only copied and pasted the logic. Need to properly customize it for Java

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

func AddSpringBootTemplate(d DomainTemplate) error {
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
