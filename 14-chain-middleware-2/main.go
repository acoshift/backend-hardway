package main

import (
	"log"
	"net/http"
)

func main() {
	h := chain(
		m1,
		m2,
	)(http.HandlerFunc(index))

	http.ListenAndServe(":3333", h)
}

type middleware func(http.Handler) http.Handler

func chain(ms ...middleware) middleware {
	// implement
	return nil
}

func m1(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("m1 called")
		h.ServeHTTP(w, r)
	})
}

func m2(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("m2 called")
		h.ServeHTTP(w, r)
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index"))
}
