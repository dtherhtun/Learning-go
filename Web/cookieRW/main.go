package main

import (
	"fmt"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some valueeee",
		Path:  "/",
	})
	fmt.Fprintln(w, "Check the Browser")
}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}
	fmt.Fprintln(w, "Here is your COOKIE: ", c)
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
