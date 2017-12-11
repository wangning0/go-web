package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9002",
	}

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.URL.Query())
		fmt.Fprint(w, r.URL.RawQuery, r.URL.Path)
	})

	http.HandleFunc("/headers", func (w http.ResponseWriter, r *http.Request) {
		h := r.Header
		fmt.Fprintln(w, h.Get("Accept-Encoding"))
	})

	http.HandleFunc("/body", func (w http.ResponseWriter, r *http.Request) {
		len := r.ContentLength
		body := make([]byte, len)
		r.Body.Read(body)
		fmt.Fprintln(w, string(body))
	})

	server.ListenAndServe()
}