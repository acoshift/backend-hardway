package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3333", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// read incoming cookie
	i := 0
	if c, _ := r.Cookie("data"); c != nil {
		i, _ = strconv.Atoi(c.Value)
	}
	i++

	// set new cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    strconv.Itoa(i),
		Path:     "/",
		HttpOnly: true,
		MaxAge:   5,
	})

	fmt.Fprintf(w, "new cookie: %d", i)
}
