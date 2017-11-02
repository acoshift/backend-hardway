package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":3333", http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Content-Type: multipart/form-data

	// parse multipart form

	// print form values and files to console

	w.Write([]byte("ok\n"))
}
