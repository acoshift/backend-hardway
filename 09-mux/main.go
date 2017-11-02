package main

import "net/http"

func main() {
	// create *http.ServeMux with http.NewServeMux()

	// add handlers to mux

	// create http server
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about"))
}
