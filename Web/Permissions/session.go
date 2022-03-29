package main

import "net/http"

var dbSessions = map[string]string{}
var dbUsers = map[string]user{}

func getUser(r *http.Request) user {
	var u user
	c, err := r.Cookie("session")
	if err != nil {
		return u
	}
	if email, ok := dbSessions[c.Value]; ok {
		u = dbUsers[email]
	}
	return u
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}

	email := dbSessions[c.Value]
	_, ok := dbUsers[email]
	return ok
}
