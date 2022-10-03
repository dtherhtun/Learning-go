package main

import (
	"net/http"
	"sync"
)

func newMux(todoFile string) http.Handler {
	m := http.NewServeMux()
	mu := &sync.Mutex{}

	m.HandleFunc("/", rootHandler)
	t := todoRouter(todoFile, mu)

	m.Handle("/todo", http.StripPrefix("/todo", t))
	m.Handle("/todo/", http.StripPrefix("/todo/", t))

	return m
}

func replyTextContent(w http.ResponseWriter, r *http.Request, status int, content string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	w.Write([]byte(content))
}
