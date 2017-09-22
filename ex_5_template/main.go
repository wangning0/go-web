package main

import (
	"html/template"
	"net/http"
	"fmt"
)

type ContactDetails struct {
	Name string
	Email string
	Message string
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		details := ContactDetails{
			Name: r.FormValue("name"),
			Email: r.FormValue("email"),
			Message: r.FormValue("message"),
		}
		fmt.Println(r.Form)
		fmt.Println(r.Header)
		_ = details

		tmpl.Execute(w, struct { Success bool} {true})
	}) 

	http.ListenAndServe(":8000", nil)
}