package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, ` <img src="/dog.webp">`)
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dog.webp")
}

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/dog.webp", dogPic)
	http.ListenAndServe(":8080", nil)
}
