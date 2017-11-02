package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":3333", http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// create struct

	// decode request body with io.LimitReader

	// if body too large send error to responseWriter

	// if no error, send ok to responseWriter
}
