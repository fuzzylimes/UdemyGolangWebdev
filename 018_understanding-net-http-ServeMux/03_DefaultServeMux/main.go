package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "doggy doggy doggy")
}

type hotcat int

func (m hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "kitty kitty kitty")
}

func main() {
	var d hotdog
	var c hotcat

	http.Handle("/dog/", d)
	http.Handle("/cat", c)

	http.ListenAndServe(":8080", nil) // nil is used to tell to use DefaultMux
}
