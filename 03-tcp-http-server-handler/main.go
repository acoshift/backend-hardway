package main

import (
	"net/http"
)

func main() {
	// create tcp listener at :3333
	// don't forget to close tcp listener when done

	for {
		// accept connection from listener

		go func() {
			// don't forget to close connection when done

			// create bufio.NewReader from connection

			// use http.ReadRequest to parse HTTP request

			// create new responseWriter

			// call handler with responseWriter and request
		}()
	}
}

type responseWriter struct {
	//
}

func (w *responseWriter) Header() http.Header {
	return nil
}

func (w *responseWriter) Write(p []byte) (int, error) {
	if true { // check is header not written
		// write header with status code 200 (default)
	}
	// write p to connection
	return 0, nil // return wrote bytes and/or error
}

func (w *responseWriter) WriteHeader(code int) {
	// DO NOT write header > 1 time

	// write HTTP version and status code
	// write HTTP headers
	// write empty line to start HTTP body
}

func handler(w http.ResponseWriter, r *http.Request) {
	// write data to responseWriter
}
