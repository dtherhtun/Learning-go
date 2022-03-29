package main

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Email, Fname, Lname, Role string
	Password                  []byte
}

var tpl *template.Template

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

func admin(w http.ResponseWriter, r *http.Request) {
	u := getUser(r)
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "admin" {
		http.Error(w, "You must be admin to enter the admin page", http.StatusForbidden)
		return
	}
	tpl.ExecuteTemplate(w, "admin.html", u)
}

func signUp(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		fname := r.FormValue("fname")
		lname := r.FormValue("lname")
		password := r.FormValue("password")
		role := r.FormValue("role")

		if _, ok := dbUsers[email]; ok {
			http.Error(w, "Email alrady taken!", http.StatusForbidden)
			return
		}

		c := &http.Cookie{
			Name:  "session",
			Value: uuid.NewString(),
		}
		http.SetCookie(w, c)

		bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		dbSessions[c.Value] = email
		dbUsers[email] = user{email, fname, lname, role, bs}

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
		password := r.FormValue("password")

		u, ok := dbUsers[email]
		if !ok {
			http.Error(w, "Username or Password do not match!", http.StatusForbidden)
			return
		}
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(password))
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
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/signup", signUp)
	http.HandleFunc("/login", logIn)
	http.HandleFunc("/logout", logOut)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
