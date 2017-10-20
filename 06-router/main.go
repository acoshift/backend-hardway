package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe("localhost:3333", http.HandlerFunc(router))
}

func router(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		index(w, r)
	case "/about":
		about(w, r)
	default:
		notFound(w, r)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound) // 404
	w.Write([]byte("404 page not found"))
}
