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
	// sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	sages := map[string]string{
		"India":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Muhammad",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
