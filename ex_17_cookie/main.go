package main

import (
	"net/http"
	"fmt"
	"time"
)
/*
	type Cookie struct {
		Name  		string
		Value  		string
		Path		string
		Domain		string
		Expires		time.Time
		RawExpires	string
		MaxAge		int
		Secure		bool
		HttpOnly	bool
		Raw			string
		Unparsed 	[]string //Raw text of unparsed attribute-value pairs
	}
*/
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cookie_, _:= r.Cookie("username")
		fmt.Println(cookie_)
		expiration := time.Now()
		expiration = expiration.AddDate(1, 0, 0)
		cookie := http.Cookie{Name: "username", Value: "wingerwang", Expires: expiration,}
		http.SetCookie(w, &cookie)
		fmt.Fprint(w, cookie)
	})
	http.ListenAndServe(":8000", nil)
}