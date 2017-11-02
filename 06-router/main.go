package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":3333", http.HandlerFunc(router))
}

func router(w http.ResponseWriter, r *http.Request) {
	// implement router
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
