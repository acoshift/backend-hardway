package main

import (
	"log"
	"net"
	"net/http"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:3333")
	if err != nil {
		log.Fatal("tcp listen error:", err)
	}
	defer lis.Close()

	http.Serve(lis, http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World."))
}
