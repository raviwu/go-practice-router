package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "home") }))
	http.Handle("/dog/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "doggy") }))
	http.Handle("/me/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "me, Ravi") }))
	http.ListenAndServe(":8080", nil)
}
