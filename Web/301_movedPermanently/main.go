package main

import (
	"fmt"
	"net/http"
)

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method from from bar:", r.Method)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method from foo:", r.Method)
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
