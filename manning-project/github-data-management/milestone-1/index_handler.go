package main

import (
	_ "embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed index.html.tmpl
var indexHtml string

func indexHandler(w http.ResponseWriter, req *http.Request) {
	s := initiateLoginIfRequired(w, req)
	if s == nil {
		return
	}
	indexHtmlTmpl, err := template.New("indexhtml").Parse(indexHtml)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Unexpected error", http.StatusInternalServerError)
		return
	}
	err = indexHtmlTmpl.Execute(w, sessionsStore[s.ID])
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Unexpected error", http.StatusInternalServerError)
		return
	}
}
