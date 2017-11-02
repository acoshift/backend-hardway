package main

import (
	"net/http"
)

func main() {
	// start web page on another port
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`
				<!doctype html>
				<button onclick="invokeApi('/with-cors')">Fetch API with CORS</button>
				<button onclick="invokeApi('/no-cors')">Fetch API without CORS</button>
				<div id=result></div>
				<script>
					function invokeApi (path) {
						const result = document.querySelector('#result')
						result.innerHTML = ''

						fetch('http://localhost:3333' + path)
							.then((resp) => {
								result.innerHTML += 'X-Request-Id: ' + resp.headers.get('X-Request-Id') + '<br>'
								return resp.text()
							})
							.then((res) => {
								result.innerHTML += res
							})
							.catch((err) => {
								result.innerHTML += err
							})
					}
				</script>
			`))
		})

		http.ListenAndServe(":8080", mux)
	}()

	mux := http.NewServeMux()
	mux.Handle("/with-cors", cors(http.HandlerFunc(result)))
	mux.Handle("/no-cors", http.HandlerFunc(result))

	http.ListenAndServe(":3333", mux)
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// only allow http://localhost:8080
		if r.Method == http.MethodOptions {
			// send forbidden if origin not allowed

			// set preflight headers

			// write header
			return
		}

		// set real headers

		// toggle Access-Control-Expose-Headers to see result in browser

		h.ServeHTTP(w, r)
	})
}

func result(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Request-Id", "1234")
	w.Write([]byte(`{"name":"launcher-1234"}`))
}
