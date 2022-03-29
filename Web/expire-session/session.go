package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

var dbUsers = map[string]user{}
var dbSessions = map[string]session{}
var dbSessionsCleaned time.Time

const sessionLength int = 30

func getUser(w http.ResponseWriter, r *http.Request) user {
	c, err := r.Cookie("session")
	if err != nil {
		c = &http.Cookie{
			Name:  "session",
			Value: uuid.NewString(),
		}
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	var u user
	if s, ok := dbSessions[c.Value]; ok {
		s.LastActivity = time.Now()
		dbSessions[c.Value] = s
		u = dbUsers[s.User]
	}
	return u
}

func alreadyLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := dbSessions[c.Value]
	if ok {
		s.LastActivity = time.Now()
		dbSessions[c.Value] = s
	}
	_, ok = dbUsers[s.User]
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return ok
}

func cleanSessions() {
	fmt.Println("Befor Clean")
	showSessions()
	for k, v := range dbSessions {
		if time.Since(v.LastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}
}

func showSessions() {
	fmt.Println("**********")
	for k, v := range dbSessions {
		fmt.Println(k, v.User)
	}
	fmt.Println("")
}
