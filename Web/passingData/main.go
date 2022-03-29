package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func foo(w http.ResponseWriter, r *http.Request) {
	f := r.FormValue("first")
	l := r.FormValue("last")
	s := r.FormValue("subscribe") == "on"

	err := tpl.ExecuteTemplate(w, "index.html", person{f, l, s})
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
