package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type user struct {
	Email, First, Last string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func foo(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")

	if err == http.ErrNoCookie {

		c = &http.Cookie{
			Name:  "session",
			Value: uuid.NewString(),
		}
		http.SetCookie(w, c)
	}

	var u user

	if e, ok := dbSessions[c.Value]; ok {
		u = dbUsers[e]
	}

	if r.Method == http.MethodPost {
		e := r.FormValue("email")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		u = user{e, f, l}
		dbSessions[c.Value] = e
		dbUsers[e] = u
	}

	err = tpl.ExecuteTemplate(w, "index.html", u)
	if err != nil {
		log.Fatalln("Error execute index template :", err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	e, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	err = tpl.ExecuteTemplate(w, "bar.html", dbUsers[e])
	if err != nil {
		log.Fatalln("Error execute bar template :", err)
	}
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)
}
