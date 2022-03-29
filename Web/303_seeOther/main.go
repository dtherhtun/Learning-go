package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at bar:", r.Method)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func barred(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at barrred", r.Method)
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at foo", r.Method)
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
