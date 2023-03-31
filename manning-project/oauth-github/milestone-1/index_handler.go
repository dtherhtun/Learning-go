package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	s, err := getSession(r)
	if err != nil {
		stateToken, err := getRandomString()
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		githubLoginUrl := oauthConf.AuthCodeURL(stateToken)
		setCookie(w, oauthStateCookie, stateToken, 600)
		http.Redirect(w, r, githubLoginUrl, http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(w, "Successfully authorized to access GitHub on your behalf: %#v", sessionsStore[s.ID].Login)
}
