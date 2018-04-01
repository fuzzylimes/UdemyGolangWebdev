package main

import (
	"io"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", count)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func count(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("visits")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "visits",
			Value: "0",
		}
	}

	count, _ := strconv.Atoi(cookie.Value)
	count++
	cookie.Value = strconv.Itoa(count)
	http.SetCookie(w, cookie)
	io.WriteString(w, cookie.Value)

}
