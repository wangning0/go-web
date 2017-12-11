package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct {}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO!")
}

func protect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ...
		h.ServeHTTP(w, r)
	})
}
// HandlerFunc(f)通过调用f实现了Handler接口
func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		fmt.Printf("called")
		h.ServeHTTP(w, r)
	})
}
func main() {
	server := http.Server {
		Addr: "127.0.0.1:9001",
	}
	hello := HelloHandler{}
	http.Handle("/hello", protect(log(hello)))
}