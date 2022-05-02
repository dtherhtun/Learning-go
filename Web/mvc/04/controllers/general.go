package controllers

import (
	"net/http"
	"text/template"

	"github.com/dtherhtun/Learning-go/Web/mvc/04/session"
)

type Controller struct {
	tpl *template.Template
}

func NewController(t *template.Template) *Controller {
	return &Controller{t}
}

func (c Controller) Index(w http.ResponseWriter, r *http.Request) {
	u := session.GetUser(w, r)
	session.ShowSession()
	c.tpl.ExecuteTemplate(w, "index.html", u)
}

func (c Controller) Bar(w http.ResponseWriter, r *http.Request) {
	u := session.GetUser(w, r)
	if !session.AlreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "admin" {
		http.Error(w, "You must be admin to login this page", http.StatusForbidden)
		return
	}
	session.ShowSession()
	c.tpl.ExecuteTemplate(w, "bar.html", u)
}
