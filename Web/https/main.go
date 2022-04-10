package main

import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain")
	w.Write([]byte("This is an example Server.\n"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}

// Generate unsigned certificate
// go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=somedomainname.com
// for example
// go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost
