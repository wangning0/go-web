/*
	关于json数据的处理
	使用encoding/json 这个包处理
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Age  	  int	   `json:"age"` 
}

func main() {
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
	})

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		peter := User {
			Firstname: "winger",
			Lastname: "wang",
			Age: 21,
		}
		json.NewEncoder(w).Encode(peter)
	})

	http.ListenAndServe(":8080", nil)
}