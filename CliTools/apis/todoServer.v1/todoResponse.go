package main

import (
	"encoding/json"
	"github.com/dtherhtun/Learning-go/CliTools/interacting/todo"
	"log"
	"net/http"
	"time"
)

type todoResponse struct {
	Results todo.List `json:"results"`
}

func (r todoResponse) MarshalJSON() ([]byte, error) {
	resp := struct {
		Results      todo.List `json:"results"`
		Date         int64     `json:"date"`
		TotalResults int       `json:"total_results"`
	}{
		Results:      r.Results,
		Date:         time.Now().Unix(),
		TotalResults: len(r.Results),
	}

	return json.Marshal(resp)
}

func replyJSONContent(w http.ResponseWriter, r *http.Request, status int, resp *todoResponse) {

	body, err := json.Marshal(resp)
	if err != nil {
		replyError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(body)
}

func replyError(w http.ResponseWriter, r *http.Request, status int, message string) {
	log.Printf("%s %s: Error: %d %s", r.URL, r.Method, status, message)
	http.Error(w, http.StatusText(status), status)
}
