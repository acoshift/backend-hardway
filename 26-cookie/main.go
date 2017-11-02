package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3333", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	i := 0

	// read i from cookie named `data`

	i++

	// set i to new cookie named `data`

	fmt.Fprintf(w, "new cookie: %d", i)
}
