package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.ListenAndServe(":3333", http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Content-Type: application/json

	var v struct {
		Name  string `json:"name"`
		Score int    `json:"score"`
	}

	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		log.Println("decode error:", err)
		return
	}
	log.Println(v)
	w.Write([]byte("ok\n"))
}
