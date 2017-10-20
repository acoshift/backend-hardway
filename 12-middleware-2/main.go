package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	r := router{}
	r.Get("/", http.HandlerFunc(index))
	r.Get("/about", http.HandlerFunc(about))

	h := logger(&r)

	http.ListenAndServe("localhost:3333", h)
}

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nw := &responseWriter{ResponseWriter: w}
		start := time.Now()
		h.ServeHTTP(nw, r)
		log.Printf("[%d] %s %s %v", nw.code, r.Method, r.RequestURI, time.Now().Sub(start))
	})
}

type responseWriter struct {
	http.ResponseWriter
	code        int
	wroteHeader bool
}

func (w *responseWriter) WriteHeader(code int) {
	if w.wroteHeader {
		return
	}
	w.wroteHeader = true
	w.code = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *responseWriter) Write(p []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	return w.ResponseWriter.Write(p)
}

type router struct {
	// path => method => handler
	path map[string]map[string]http.Handler
}

type path struct {
	Method  string
	Path    string
	Handler http.Handler
}

func (router *router) Add(m, p string, h http.Handler) {
	if router.path == nil {
		router.path = make(map[string]map[string]http.Handler)
	}

	if router.path[p] == nil {
		router.path[p] = make(map[string]http.Handler)
	}

	router.path[p][m] = h
}

func (router *router) Get(p string, h http.Handler) {
	router.Add(http.MethodGet, p, h)
}

func (router *router) Post(p string, h http.Handler) {
	router.Add(http.MethodPost, p, h)
}

func (router *router) Put(p string, h http.Handler) {
	router.Add(http.MethodPut, p, h)
}

func (router *router) Patch(p string, h http.Handler) {
	router.Add(http.MethodPatch, p, h)
}

func (router *router) Delete(p string, h http.Handler) {
	router.Add(http.MethodDelete, p, h)
}

func (router *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if router.path == nil {
		http.NotFound(w, r)
		return
	}

	p := router.path[r.URL.Path]
	if p == nil {
		http.NotFound(w, r)
		return
	}

	h := p[r.Method]
	if h == nil {
		http.NotFound(w, r)
		return
	}

	h.ServeHTTP(w, r)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about"))
}
