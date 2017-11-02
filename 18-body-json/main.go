package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":3333", http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Content-Type: application/json

	// create request body struct

	// decode request body using json.NewDecoder

	// print decoded struct to console

	w.Write([]byte("ok\n"))
}
