package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello dog")
}

func c(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello cat")
}

func main() {
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)
}
