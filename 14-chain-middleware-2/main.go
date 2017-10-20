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

	http.ListenAndServe("localhost:3333", h)
}

type middleware func(http.Handler) http.Handler

func chain(ms ...middleware) middleware {
	return func(h http.Handler) http.Handler {
		for i := len(ms); i > 0; i-- {
			h = ms[i-1](h)
		}
		return h
	}
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
