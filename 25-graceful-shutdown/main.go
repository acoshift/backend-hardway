package main

import (
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", index)

	// create http.Server

	// start server on another goroutine

	// create buffered (size=1) channel for os.Signal

	// call signal.Notify to notify channel when received syscall.SIGTERM

	// block until receive signal

	// shutdown server
}

func index(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	w.Write([]byte("ok"))
}
