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

	data := struct {
		Words []string
		Lname string
	}{
		xs,
		"limes",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", data)
	if err != nil {
		log.Fatalln(err)
	}
}
