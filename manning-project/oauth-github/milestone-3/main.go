package main

import (
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var (
	oauthConf       *oauth2.Config
	oauthHttpClient *http.Client
)

func initOAuthConfig(getEnvironValue func(string) string) {

	if len(getEnvironValue("CLIENT_ID")) == 0 || len(getEnvironValue("CLIENT_SECRET")) == 0 {
		log.Fatal("Must specify your app's CLIENT_ID and CLIENT_SECRET")
	}

	oauthConf = &oauth2.Config{
		ClientID:     getEnvironValue("CLIENT_ID"),
		ClientSecret: getEnvironValue("CLIENT_SECRET"),
		Scopes:       []string{"repo", "user"}, // see the project desrciption for understandng why we need full scopes here
		Endpoint:     github.Endpoint,
	}
}

func registerHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/github/callback", githubCallbackHandler)
}

func main() {

	initOAuthConfig(os.Getenv)

	mux := http.NewServeMux()
	registerHandlers(mux)
	log.Println("Starting server..")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
