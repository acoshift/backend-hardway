package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", resizeHandler)
	http.ListenAndServe(":3333", nil)
}

func resizeHandler(w http.ResponseWriter, r *http.Request) {
	// DO NOT accept other method than POST

	// limit body to 2 MiB

	// decode image using image.Decode

	// print image type (jpg, png, bmp) to console

	// use imaging.Thumbnail to resize image

	// encode result image using jpeg.Encode

	// write result jpeg image to responseWriter
}
