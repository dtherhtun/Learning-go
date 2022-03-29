package main

import (
	"fmt"
	"io"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/dog":
		io.WriteString(w, "helo dog")
	case "/cat":
		fmt.Fprintln(w, "hello cat")
	}
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}
