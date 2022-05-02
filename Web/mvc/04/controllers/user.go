package controllers

import (
	"net/http"
	"time"

	"github.com/dtherhtun/Learning-Go/Web/mvc/04/models"
	"github.com/dtherhtun/Learning-go/Web/mvc/04/session"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (c Controller) SignUp(w http.ResponseWriter, r *http.Request) {
	if !session.AlreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var u models.User

	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		pw := r.FormValue("password")
		name := r.FormValue("name")
		role := r.FormValue("role")
		gen := r.FormValue("gender")
		age := r.FormValue("age")

		if _, ok := session.Users[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		ck := &http.Cookie{
			Name:  "session",
			Value: uuid.New().String(),
		}
		ck.MaxAge = session.Length
		http.SetCookie(w, ck)
		session.Sessions[ck.Value] = models.Session{un, time.Now()}
		bs, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		u = models.User{un, bs, name, role, gen, age}
		session.Users[un] = u
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}
