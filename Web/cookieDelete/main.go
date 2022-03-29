package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello Earth</h1>")
	fmt.Fprintln(w, `<h2><a href="/set">Set cookie</a></h2>`)
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: "we are fucking stoner",
	})
	fmt.Fprintln(w, "<h1> Cookie has been setup</h1>")
	fmt.Fprintln(w, `<h2><a href="/read">Read cookie</a></h2>`)
}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}
	fmt.Fprintf(w, `<h1> Here is your cookie: %v</h1><br><h2><a href="/expire">Expire</a></h2>`, c)
}

func expire(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}
	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":8080", nil)
}
