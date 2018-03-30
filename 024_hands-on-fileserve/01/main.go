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
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dogPic", dogPic)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "toby.jpg")
}
