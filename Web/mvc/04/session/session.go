package session

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dtherhtun/Learning-Go/Web/mvc/04/models"
	"github.com/google/uuid"
)

const Length int = 30

var Users = map[string]models.User{}
var Sessions = map[string]models.Session{}

func GetUser(w http.ResponseWriter, r *http.Request) models.User {
	ck, err := r.Cookie("session")
	if err != nil {
		ck = &http.Cookie{
			Name:  "session",
			Value: uuid.New().String(),
		}
	}
	ck.MaxAge = Length
	http.SetCookie(w, ck)

	u := models.User{}

	if s, ok := Sessions[ck.Value]; ok {
		s.LastActivity = time.Now()
		Sessions[ck.Value] = s
		u = Users[s.UserName]
	}
	return u
}

func AlreadyLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	ck, err := r.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := Sessions[ck.Value]
	if ok {
		s.LastActivity = time.Now()
		Sessions[ck.Value] = s
	}
	_, ok = Users[s.UserName]

	ck.MaxAge = Length
	http.SetCookie(w, ck)
	return ok
}

func CleanSession() {
	fmt.Println("Before Clean")
	ShowSession()
	for k, v := range Sessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 30) {
			delete(Sessions, k)
		}
	}
}

func ShowSession() {
	fmt.Println("***********")
	for k, v := range Sessions {
		fmt.Println(k, v.UserName)
	}
	fmt.Println()
}
