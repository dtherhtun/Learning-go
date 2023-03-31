package main

import (
	"context"
	"net/http"

	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

func githubCallbackHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if !validCallback(r) {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	code := r.URL.Query().Get("code")
	ctx = context.WithValue(ctx, oauth2.HTTPClient, oauthHttpClient)
	token, err := oauthConf.Exchange(ctx, code)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}

	client := github.NewClient(oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token.AccessToken},
	)))

	user, _, err := client.Users.Get(ctx, "")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	sessionID, _ := getRandomString()

	sessionsStore[sessionID] = userData{Login: *user.Login, accessToken: token.AccessToken}

	setCookie(w, sessionCookie, sessionID, sessionCookieMaxAge)
	setCookie(w, oauthStateCookie, "", -1)

	http.Redirect(w, r, "/", http.StatusFound)
}
