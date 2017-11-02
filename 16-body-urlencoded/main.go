package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":3333", http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Content-Type: application/x-www-form-urlencoded

	// print request body to console

	// use url.ParseQuery to parse x-www-form-urlencoded

	// print result to console

	w.Write([]byte("ok\n"))
}
