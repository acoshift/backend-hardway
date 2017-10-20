package main

import (
	"bufio"
	"fmt"
	"io"
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

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal("tcp accept error:", err)
		}

		go func() {
			defer conn.Close()
			reader := bufio.NewReader(conn)
			r, err := http.ReadRequest(reader)
			if err != nil {
				log.Printf("read request: %v\n", err)
				return
			}
			w := &responseWriter{writer: conn}
			handler(w, r)
		}()
	}
}

type responseWriter struct {
	header      http.Header
	wroteHeader bool
	writer      io.Writer
}

func (w *responseWriter) Header() http.Header {
	if w.header == nil {
		w.header = make(http.Header)
	}
	return w.header
}

func (w *responseWriter) Write(p []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(200)
	}
	return w.writer.Write(p)
}

func (w *responseWriter) WriteHeader(code int) {
	if w.wroteHeader {
		return
	}
	w.wroteHeader = true
	fmt.Fprintf(w.writer, "HTTP/1.1 %d %s\n", code, http.StatusText(code))
	w.header.Write(w.writer)
	w.writer.Write([]byte("\n"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World."))
}
