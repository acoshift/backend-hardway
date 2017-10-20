package main

import (
	"image"
	"image/jpeg"
	"io"
	"net/http"

	"github.com/disintegration/imaging"
)

func main() {
	http.Handle("/", &resizeHandler{resizer: resizeImage})
	http.ListenAndServe(":3333", nil)
}

type resizeHandler struct {
	resizer func(io.Writer, io.Reader) error
}

func (h *resizeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Write([]byte("Allow only POST"))
		return
	}

	err := h.resizer(w, io.LimitReader(r.Body, 2<<20)) // 2 MiB
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func resizeImage(dst io.Writer, src io.Reader) error {
	img, _, err := image.Decode(src)
	if err != nil {
		return err
	}
	result := imaging.Thumbnail(img, 150, 150, imaging.Lanczos)
	err = jpeg.Encode(dst, result, &jpeg.Options{Quality: 80})
	if err != nil {
		return err
	}
	return nil
}
