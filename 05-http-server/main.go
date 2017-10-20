package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe("localhost:3333", http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World."))
}
