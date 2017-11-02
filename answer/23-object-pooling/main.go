package main

import (
	"compress/gzip"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/large", large)
	http.ListenAndServe(":3333", gzipMiddleware(mux))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	data := []byte(`
		<!doctype html>
		<title>Compression Test</title>
		<h1>Index</h1>
		<p>Hello, World!</p>
	`)
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Write(data)
}

func large(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	data := []byte(`
		<!doctype html>
		<title>Compression Test</title>
		<h1>Large</h1>
		<p>The quick, brown fox jumps over a lazy dog. DJs flock by when MTV ax quiz prog. Junk MTV quiz graced by fox whelps. Bawds jog, flick quartz, vex nymphs. Waltz, bad nymph, for quick jigs vex! Fox nymphs grab quick-jived waltz. Brick quiz whangs jumpy veldt fox. Bright vixens jump; dozy fowl quack. Quick wafting zephyrs vex bold Jim. Quick zephyrs blow, vexing daft Jim. Sex-charged fop blew my junk TV quiz. How quickly daft jumping zebras vex. Two driven jocks help fax my big quiz. Quick, Baz, get my woven flax jodhpurs! "Now fax quiz Jack!" my brave ghost pled. Five quacking zephyrs jolt my wax bed. Flummoxed by job, kvetching W. zaps Iraq. Cozy sphinx waves quart jug of bad milk. A very bad quack might jinx zippy fowls. Few quips galvanized the mock jury box. Quick brown dogs jump over the lazy fox. The jay, pig, fox, zebra, and my wolves quack! Blowzy red vixens fight for a quick jump. Joaquin Phoenix was gazed by MTV for luck. A wizard’s job is to vex chumps quickly in fog. Watch "Jeopardy!", Alex Trebek's fun TV quiz game. Woven silk pyjamas exchanged for blue quartz. Brawny gods just</p>
		<p>Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi. Aenean vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui. Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem. Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna. Sed consequat, leo eget bibendum sodales, augue velit cursus nunc,</p>
	`)
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func gzipMiddleware(h http.Handler) http.Handler {
	pool := &sync.Pool{
		New: func() interface{} {
			return gzip.NewWriter(ioutil.Discard)
		},
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			h.ServeHTTP(w, r)
			return
		}

		if len(r.Header.Get("Sec-WebSocket-Key")) > 0 {
			h.ServeHTTP(w, r)
			return
		}

		wh := w.Header()

		if len(wh.Get("Content-Encoding")) > 0 {
			h.ServeHTTP(w, r)
			return
		}

		wh.Set("Vary", "Accept-Encoding")

		gw := &gzipResponseWriter{
			ResponseWriter: w,
			pool:           pool,
		}
		defer gw.Close()

		h.ServeHTTP(gw, r)
	})
}

type gzipResponseWriter struct {
	http.ResponseWriter
	pool          *sync.Pool
	gzipWriter    *gzip.Writer
	contentLength int
}

func (w *gzipResponseWriter) init() {
	h := w.Header()

	if len(h.Get("Content-Encoding")) > 0 {
		return
	}

	if w.contentLength == 0 {
		w.contentLength, _ = strconv.Atoi(h.Get("Content-Length"))
	}
	if w.contentLength > 0 && w.contentLength <= 860 {
		return
	}

	w.gzipWriter = w.pool.Get().(*gzip.Writer)
	w.gzipWriter.Reset(w.ResponseWriter)
	h.Del("Content-Length")
	h.Set("Content-Encoding", "gzip")
}

func (w *gzipResponseWriter) Write(p []byte) (int, error) {
	if w.gzipWriter == nil {
		w.init()
	}

	if w.gzipWriter != nil {
		return w.gzipWriter.Write(p)
	}

	return w.ResponseWriter.Write(p)
}

func (w *gzipResponseWriter) Close() {
	if w.gzipWriter == nil {
		return
	}
	w.gzipWriter.Close()
	w.pool.Put(w.gzipWriter)
}

func (w *gzipResponseWriter) WriteHeader(code int) {
	if w.gzipWriter == nil {
		w.init()
	}
	w.ResponseWriter.WriteHeader(code)
}
