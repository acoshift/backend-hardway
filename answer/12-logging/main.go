package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/about", about)

	h := logger(mux)

	http.ListenAndServe(":3333", h)
}

type logRecord struct {
	Time         string `json:"time"`
	RemoteIP     string `json:"remote_ip"`
	Host         string `json:"host"`
	Method       string `json:"method"`
	URI          string `json:"uri"`
	Status       int    `json:"status"`
	Latency      int64  `json:"latency"`
	LatencyHuman string `json:"latency_human"`
	BytesIn      int64  `json:"bytes_in"`
	BytesOut     int64  `json:"bytes_out"`
}

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		record := logRecord{
			Time:     start.UTC().Format(time.RFC3339Nano),
			RemoteIP: r.RemoteAddr,
			Host:     r.Host,
			Method:   r.Method,
			URI:      r.RequestURI,
			BytesIn:  r.ContentLength,
		}

		nw := &responseWriter{ResponseWriter: w}
		h.ServeHTTP(nw, r)

		diff := time.Now().Sub(start)
		record.Latency = int64(diff)
		record.LatencyHuman = diff.String()
		record.Status = nw.code
		record.BytesOut = nw.wroteLenght

		logStr, _ := json.Marshal(&record)
		fmt.Println(string(logStr))
	})
}

type responseWriter struct {
	http.ResponseWriter
	code        int
	wroteHeader bool
	wroteLenght int64
}

func (w *responseWriter) WriteHeader(code int) {
	if w.wroteHeader {
		return
	}
	w.wroteHeader = true
	w.code = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *responseWriter) Write(p []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	l, err := w.ResponseWriter.Write(p)
	w.wroteLenght += int64(l)
	return l, err
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about"))
}
