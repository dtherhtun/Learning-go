package main

import (
	"fmt"
	"net/http"
)

type hotdog string

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("DTher-key", "fuckyou")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1> Any Code you want to this function</h1>")
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}
