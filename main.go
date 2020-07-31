package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "home") })
	http.HandleFunc("/dog/", func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "doggy") })
	http.HandleFunc("/me/", func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "me, Ravi") })
	http.ListenAndServe(":8080", nil)
}
