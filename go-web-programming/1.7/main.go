/*
type Cookie struct {
	Name string 
	Value string
	Path string
	Domain string
	Expires time.Time
	RawExpires string
	MaxAge int
	Secure bool
	HttpOnly bool
	Raw string
	Unparsed []string
}
**/

package main

import (
	"net/http"
	"fmt"
)
// func setCookie (w http.ResponseWriter, r *http.Request) {
	// c1 := http.Cookie {
	// 	Name: "first_cookie",
	// 	Value: "GWP",
	// 	HttpOnly: true,
	// }
	// c2 := http.Cookie {
	// 	Name: "second_cookie",
	// 	Value: "cookie_",
	// 	HttpOnly: true,
	// }
// 	w.Header().Set("Set-Cookie", c1.String())
// 	w.Header().Add("Set-Cookie", c2.String())
// }
func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie {
		Name: "first_cookie",
		Value: "GWP",
		HttpOnly: true,
	}
	c2 := http.Cookie {
		Name: "second_cookie",
		Value: "cookie_",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Cookie"]
	fmt.Fprintln(w, h)
}
func otherGetCookie(w http.ResponseWriter, r *http.Request) {
	cl, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "cannot get first cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, cl)
	fmt.Fprintln(w, cs)
}
func main() {
	server := http.Server {
		Addr: "127.0.0.1:8000",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	http.HandleFunc("/ç", otherGetCookie)
	server.ListenAndServe()
}