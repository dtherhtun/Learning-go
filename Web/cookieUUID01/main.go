package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func foo(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err == http.ErrNoCookie {
		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    uuid.NewString(),
			HttpOnly: true,
		})
	}
	fmt.Println(c)
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favico.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
