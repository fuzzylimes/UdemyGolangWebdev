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

	xs := []string{"zero", "one", "two", "three", "four", "five"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", xs)
	if err != nil {
		log.Fatalln(err)
	}
}
