package main

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Email, Fname, Lname string
	Pass                []byte
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
		passwd := r.FormValue("passwd")

		if _, ok := dbUsers[email]; ok {
			http.Error(w, "Email alrady taken", http.StatusForbidden)
			return
		}

		c := &http.Cookie{
			Name:  "session",
			Value: uuid.NewString(),
		}
		http.SetCookie(w, c)
		bs, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal Server Error!", http.StatusInternalServerError)
			return
		}
		u = user{email, fname, lname, bs}
		dbSessions[c.Value] = email
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
		passwd := r.FormValue("passwd")

		u, ok := dbUsers[email]
		if !ok {
			http.Error(w, "Username or password do not match!", http.StatusForbidden)
			return
		}
		err := bcrypt.CompareHashAndPassword(u.Pass, []byte(passwd))
		if err != nil {
			http.Error(w, "Username or Password do not match!", http.StatusForbidden)
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

func logOut(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	c, _ := r.Cookie("session")
	delete(dbSessions, c.Value)
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/login", logIn)
	http.HandleFunc("/logout", logOut)
	http.HandleFunc("/signup", signUp)
	http.Handle("favico.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
