package main

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Email, First, Last string
	Pass               []byte
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	bs, _ := bcrypt.GenerateFromPassword([]byte("pass99word"), bcrypt.MinCost)
	dbUsers["d@mail.com"] = user{"d@mail.com", "dther", "htun", bs}
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

	var u user

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		fname := r.FormValue("fname")
		lname := r.FormValue("lname")
		pass := r.FormValue("passwd")

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

		bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		u = user{email, fname, lname, bs}
		dbUsers[email] = u

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.html", nil)
}

func logIn(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		pass := r.FormValue("passwd")

		u, ok := dbUsers[email]
		if !ok {
			http.Error(w, "Username or Password do not match", http.StatusForbidden)
			return
		}
		err := bcrypt.CompareHashAndPassword(u.Pass, []byte(pass))
		if err != nil {
			http.Error(w, "Username or Password do not match", http.StatusForbidden)
			return
		}
		c := &http.Cookie{
			Name:  "session",
			Value: uuid.NewString(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = email
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/login", logIn)
	http.HandleFunc("/signup", signUp)
	http.Handle("/favico.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
