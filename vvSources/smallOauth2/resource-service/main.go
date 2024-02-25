package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
)

const authServiceURL = "http://auth-service:9096"

func main() {
	manager := manage.NewDefaultManager()
	// No need to configure token storage in the protected-service

	srv := server.NewDefaultServer(manager)
	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	http.HandleFunc("/protected", validateToken(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, I'm protected"))
	}, srv))

	log.Fatal(http.ListenAndServe(":9097", nil))
}

func validateToken(f http.HandlerFunc, srv *server.Server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract token from query parameter
		token := r.URL.Query().Get("access_token")

		// Validate token by making a request to auth-service's introspection endpoint
		if err := validateTokenWithAuthService(token); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		f.ServeHTTP(w, r)
	})
}

func validateTokenWithAuthService(token string) error {
	// Make a request to auth-service's introspection endpoint to validate the token
	resp, err := http.Post(authServiceURL+"/introspect", "application/x-www-form-urlencoded", strings.NewReader("token="+token))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.ErrAccessDenied
	}

	return nil
}
