package main

import (
	"os"
	"text/template"
)

type Inventory struct {
	Material string
	Count uint
}

func main() {
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} itmes are made of {{.Material}}")
	if err != nil { panic(err) }
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil { panic(err) }
}
