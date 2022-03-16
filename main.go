package main

import (
	"os"
	"text/template"
)

type model struct {
	name string
}

func main() {
	t, _ := template.ParseFiles("model-template")
	t.Execute(os.Stdout, "myapp")
}
