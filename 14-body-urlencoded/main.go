package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	http.ListenAndServe(":3333", http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	buf := bytes.Buffer{}
	io.Copy(&buf, r.Body)
	log.Println(buf.String())
	body, err := url.ParseQuery(buf.String())
	if err != nil {
		log.Println("parse error:", err)
		return
	}
	log.Println(body)
	w.Write([]byte("ok\n"))
}
