package main

import (
	"net/http"
	"html/template"
	"math/rand"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func iterator(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("iterator.html")
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t.Execute(w, daysOfWeek)
}

func formatDate(t time.Time) string {
	layout := "2016-12-02"
	return t.Format(layout)
}

func customFunc(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"fdate": formatDate,
	}
	t := template.New("tmpl.html").Funcs(funcMap)
	t, _ = t.ParseFiles("tmpl.html")
	t.Execute(w, time.Now())
}

func main() {
	server := http.Server {
		Addr: "127.0.0.1:9000",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/iterator", iterator)
	http.HandleFunc("/custom", customFunc)
	server.ListenAndServe()
}