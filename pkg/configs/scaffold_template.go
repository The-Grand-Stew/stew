package configs

import (
	"strings"
	"text/template"
)

type ScaffoldLocations struct {
	Source      string
	Destination string
	Filename    string
}

var FuncMap = template.FuncMap{
	"ToUpper": strings.ToUpper,
	"ToLower": strings.ToLower,
	"Title":   strings.Title,
	"Join":    strings.Join,
}
