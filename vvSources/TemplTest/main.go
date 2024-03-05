package main

import (
	"context"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := hello("John")
	http.Handle("/", templ.Handler(component))
	if err := http.ListenAndServe(":3333", nil); err != nil {
		log.Fatalln(err)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	component := hello("John")
	component.Render(context.Background(), w)
}
