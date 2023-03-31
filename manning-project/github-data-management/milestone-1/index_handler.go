package main

import (
	_ "embed"
	"html/template"
	"net/http"
)

//go:embed index.html.tmpl
var indexHtml string

func indexHandler(w http.ResponseWriter, req *http.Request) {
	s, err := getSession(req)
	if err != nil {
		stateToken, err := getRandomString()
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		githubLoginUrl := oauthConf.AuthCodeURL(stateToken)
		setCookie(w, oauthStateCookie, stateToken, 600)
		http.Redirect(w, req, githubLoginUrl, http.StatusTemporaryRedirect)
		return
	}
	tpl, err := template.New("index").Parse(indexHtml)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	tpl.Execute(w, sessionsStore[s.ID])
}
