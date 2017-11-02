package main

import (
	"net/http"
)

func main() {
	// create router

	// use http.ListenAndServe to create http server
}

type router struct {
	paths []*path
}

type path struct {
	//
}

func (router *router) Get(p string, h http.Handler) {
	// implement
}

func (router *router) Post(p string, h http.Handler) {
	// implement
}

func (router *router) Put(p string, h http.Handler) {
	// implement
}

func (router *router) Patch(p string, h http.Handler) {
	// implement
}

func (router *router) Delete(p string, h http.Handler) {
	// implement
}

func (router *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// implement router
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about"))
}
