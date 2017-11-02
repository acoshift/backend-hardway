package main

import (
	"log"
	"net/http"
)

func main() {
	http.ListenAndServe(":3333", http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Content-Type: multipart/form-data
	r.ParseMultipartForm(10 << 20) // 10 * 2^20
	log.Println(r.MultipartForm.Value)
	log.Println(r.MultipartForm.File)
	w.Write([]byte("ok\n"))
}
