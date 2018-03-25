package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./*.html"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", `"This is pretty great; don't you think?"`)
	if err != nil {
		log.Fatalln(err)
	}
}
