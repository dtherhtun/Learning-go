package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"
)

type githubExportResult struct {
	ID           int64    `json:"id"`
	State        string   `json:"state"`
	Repositories []string `json:"repositories"`
}

func githubNewExportApiHandler(w http.ResponseWriter, req *http.Request) {
	var respData []byte
	var repoToExportFullNames, repoExportedFullNames []string

	ctx := req.Context()

	if req.Method != http.MethodPost {
		http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
		return
	}

	s, err := getSession(req)
	if err != nil {
		http.Error(w, "session cookie invalid or not found", http.StatusUnauthorized)
		return
	}

	token := sessionsStore[s.ID].accessToken
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	ctx = context.WithValue(ctx, oauth2.HTTPClient, oauthHttpClient)
	tc := oauth2.NewClient(ctx, ts)
	ghClient := github.NewClient(tc)

	options := github.RepositoryListOptions{
		Affiliation: "owner",
		ListOptions: github.ListOptions{Page: 1},
	}
	for {
		repos, resp, err := ghClient.Repositories.List(ctx, "", &options)
		if err != nil {
			log.Println("Error retrieving user's repositories to export: ", err)
			if resp != nil {
				if respData, err := io.ReadAll(resp.Body); err == nil {
					log.Println(string(respData))
				}
				resp.Body.Close()
			}
			http.Error(w, "Error retrieving user's repositories to export", http.StatusInternalServerError)
			return
		}
		for _, repo := range repos {
			if *repo.Fork {
				continue
			}
			repoToExportFullNames = append(repoToExportFullNames, *repo.FullName)
		}

		if resp.NextPage == 0 {
			break
		}
		options.Page = resp.NextPage
	}
	m, resp, err := ghClient.Migrations.StartUserMigration(ctx, repoToExportFullNames, nil)
	if err != nil {
		log.Println("Error starting user migration: ", err)
		if resp != nil {
			if respData, err := io.ReadAll(resp.Body); err == nil {
				log.Println(string(respData))
			}
			resp.Body.Close()
		}

		http.Error(w, "Error starting user repos migration", http.StatusInternalServerError)
		return
	}

	for _, r := range m.Repositories {
		repoExportedFullNames = append(repoExportedFullNames, *r.FullName)
	}
	exportResult := githubExportResult{
		ID:           *m.ID,
		State:        *m.State,
		Repositories: repoExportedFullNames,
	}
	respData, err = json.Marshal(&exportResult)
	if err != nil {
		log.Println("Error marshalling start migration response: ", err)
		http.Error(w, "Error marshalling start migration response", http.StatusInternalServerError)
		return
	}
	w.Header().Add("content-type", "application/json")
	fmt.Fprint(w, string(respData))
}
