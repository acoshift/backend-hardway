package main

import (
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"

	"github.com/disintegration/imaging"
)

func main() {
	http.HandleFunc("/", resizeHandler)
	http.ListenAndServe(":3333", nil)
}

func resizeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Write([]byte("Allow only POST"))
		return
	}

	img, imgType, err := image.Decode(io.LimitReader(r.Body, 2<<20)) // 2 MiB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println(imgType)
	result := imaging.Thumbnail(img, 150, 150, imaging.Lanczos)
	err = jpeg.Encode(w, result, &jpeg.Options{Quality: 80})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
