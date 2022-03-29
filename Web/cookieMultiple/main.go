package main

import (
	"fmt"
	"log"
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
	c1, err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}
	fmt.Fprintln(w, "Here is your COOKIE #1: ", c1)
	c2, err := r.Cookie("general")
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintln(w, "Here is your COOKIE #2: ", c2)
	c3, err := r.Cookie("specific")
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintln(w, "Here is your COOKIE #3: ", c3)
}

func abundance(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "I love golang",
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "go love me",
	})
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
