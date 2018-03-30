package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	hello := struct {
		Name string
		Age  int
	}{
		"Bob",
		55,
	}
	err := tpl.ExecuteTemplate(w, "index.html", hello)
	HandleError(w, err)
}

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Woof")
}

func me(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "fuzzy")
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
