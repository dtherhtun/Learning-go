package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
)

const (
	oauthStateCookie    = "OAuthState"
	sessionCookie       = "Session"
	sessionCookieMaxAge = 24 * 3600 // 24 hours
)

type userData struct {
	Login       string
	accessToken string
}

var sessionsStore = make(map[string]userData)

type sessionData struct {
	ID string
}

func getRandomString() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

func validSessionID(sessionID string) bool {
	_, ok := sessionsStore[sessionID]
	return ok
}

func getSession(r *http.Request) (*sessionData, error) {
	c, err := r.Cookie(sessionCookie)
	if err != nil {
		return nil, err
	}

	if !validSessionID(c.Value) {
		return nil, fmt.Errorf("invalid session ID")
	}

	return &sessionData{ID: c.Value}, nil
}

func validCallback(r *http.Request) bool {

	gotState := r.URL.Query().Get("state")
	c, err := r.Cookie(oauthStateCookie)
	if err != nil {
		return false
	}
	if c.Value != gotState {
		return false
	}

	return true
}

func setCookie(w http.ResponseWriter, name, value string, maxAge int) {
	http.SetCookie(w, &http.Cookie{
		Name:   name,
		Value:  value,
		Path:   "/",
		MaxAge: maxAge,
	})
}
