package main

import (
	"net/http"
)

func main() {
	// create tcp listener at :3333
	// close tcp listener when done

	// use http.Serve to serve connection from listener
}

func handler(w http.ResponseWriter, r *http.Request) {
	// write data to responseWriter
}
