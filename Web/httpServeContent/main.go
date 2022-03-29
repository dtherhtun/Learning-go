package main

import (
	"io"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<image src="/dog.webp">`)
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("dog.webp")
	if err != nil {
		http.Error(w, "File Not found", http.StatusNotFound)
		return
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "File Not Found", http.StatusNotFound)
		return
	}
	http.ServeContent(w, r, f.Name(), fi.ModTime(), f)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog.webp", dogPic)
	http.ListenAndServe(":8080", nil)
}
