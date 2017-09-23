package main

import (
	"net/http"	
	"html/template"
	"strconv"
	"fmt"
	"time"
	"crypto/md5"
	"io"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method)
	if r.Method == "GET" {
		currtime := time.Now().Unix()
		fmt.Println(currtime, strconv.FormatInt(currtime, 10))
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(currtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("form.html")
		t.Execute(w, token)
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		// use token do something
		fmt.Println(token)
		fmt.Fprint(w, "token")
	}
} 

func main() {
	http.HandleFunc("/", Login)
	http.ListenAndServe(":8000", nil)
}