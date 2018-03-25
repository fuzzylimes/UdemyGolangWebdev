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

type person struct {
	Name string
	Age  int
}

func (p person) SomeProcessing() int {
	return 7
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) TakesArg(x int) int {
	return x * 2
}

func main() {

	p := person{
		"Ian Fleming",
		56,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", p)
	if err != nil {
		log.Fatalln(err)
	}
}
