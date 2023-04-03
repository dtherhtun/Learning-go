package main

import (
	_ "embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed export_new.html.tmpl
var githubNewExportHtml string

func githubNewExportViewHandler(w http.ResponseWriter, req *http.Request) {
	s := initiateLoginIfRequired(w, req)
	if s == nil {
		return
	}

	exportNewHtmlTmpl, err := template.New("github_export_view_html").Parse(githubNewExportHtml)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Unexpected error", http.StatusInternalServerError)
		return
	}
	err = exportNewHtmlTmpl.Execute(w, sessionsStore[s.ID])
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Unexpected error", http.StatusInternalServerError)
		return
	}
}
