package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Name  string
	Motto string
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseGlob("./*.html"))
}

func main() {

	u1 := user{
		"Buddah",
		"The belief of no beliefs",
		false,
	}

	u2 := user{
		"Gandhi",
		"Be the change",
		true,
	}

	u3 := user{
		"",
		"Nobody",
		true,
	}

	users := []user{u1, u2, u3}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", users)
	if err != nil {
		log.Fatalln(err)
	}
}
