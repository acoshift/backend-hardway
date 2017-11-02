package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", &resizeHandler{resizer: resizeImage})
	http.ListenAndServe(":3333", nil)
}

type resizeHandler struct {
	resizer func(io.Writer, io.Reader) error
}

func (h *resizeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// DO NOT accept other method than POST

	// limit body to 2 MiB

	// call resizer and limit body

	// write result jpeg image to responseWriter
}

func resizeImage(dst io.Writer, src io.Reader) error {
	// decode image using image.Decode

	// use imaging.Thumbnail to resize image

	// encode result image using jpeg.Encode

	// write result jpeg image to writer

	return nil
}
