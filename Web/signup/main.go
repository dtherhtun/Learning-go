package main

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

type user struct {
	Email, First, Last, Pass string
}

var tpl *template.Template
var dbSessions = map[string]string{}
var dbUsers = map[string]user{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(r)
	tpl.ExecuteTemplate(w, "index.html", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	u := getUser(r)
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "bar.html", u)
}

func signUp(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		first := r.FormValue("firstname")
		last := r.FormValue("lastname")
		pass := r.FormValue("password")

		if _, ok := dbUsers[email]; ok {
			http.Error(w, "Email already taken!", http.StatusForbidden)
			return
		}

		c := &http.Cookie{
			Name:  "session",
			Value: uuid.NewString(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = email

		u := user{email, first, last, pass}
		dbUsers[email] = u

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.html", nil)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signUp)
	http.Handle("/favico.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
