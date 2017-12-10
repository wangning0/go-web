package main

import (
	"fmt"
	"net/http"
	"data"
)

func login(writer http.ResponseWriter, request *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(writer, nil)
}

func signup(writer http.ResponseWriter, request *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "signup")
	t.Execute(writer, nil)
}

func logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	fmt.Println(cookie)
	if err != http.ErrNoCookie {
		warning(err, "Failed to get cookie")
		session := data.Session{
			Uuid: cookie.Value,
		}
		err := session.DeleteByUUID()
		if err != nil {
			danger(err, "session delete fail")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

func signup_account(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "parse form error")
	}
	user := data.User{
		Name: request.PostFormValue("name"),
		Email: request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		danger(err, "create user error")
	}
	http.Redirect(writer, request, "/login", 302)
}

func authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	fmt.Println(request.PostFormValue("email"))
	user, err := data.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		danger(err, "cannot find user")
	}
	if user.Password == data.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "cannot create session")
		}
		cookie := http.Cookie{
			Name: "_cookie",
			Value: session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
	}
}