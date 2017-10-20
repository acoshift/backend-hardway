package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func main() {
	http.ListenAndServe(":3333", http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	buf := bytes.Buffer{}
	io.Copy(&buf, r.Body)
	log.Println(buf.String())
	w.Write([]byte("ok"))
}
