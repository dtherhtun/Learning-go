package main

import (
	"io"
	"net/http"
	"os"
)

func zero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/zero.jpg">`)
}

func zeroPic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("zero.jpg")
	if err != nil {
		http.Error(w, "File Not Found", http.StatusNotFound)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}

func main() {
	http.HandleFunc("/", zero)
	http.HandleFunc("/zero.jpg", zeroPic)
	http.ListenAndServe(":8080", nil)
}
