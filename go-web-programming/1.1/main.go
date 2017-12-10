// package main

// golang handlerFunc
import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello world, %s", request.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}


// golang handle
// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// type httpServer struct {

// }

// func (server httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "hello world %s", r.URL.Path)
// }

// func main() {
// 	var server httpServer
// 	http.Handle("/", server)
// 	http.ListenAndServe(":8000", nil)
// }