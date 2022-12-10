package handlers

import (
	"log"
	"net/http"

	"github.com/dtherhtun/Learning-go/refactoring/goviolin/internal/render"
)

// Base represent the base handlers.
type Base struct {
	log *log.Logger
}

// Home handler renders the home.html page.
func (b *Base) Home(w http.ResponseWriter, r *http.Request) {
	pv := render.PageVars{
		Title: "GoViolin",
	}

	if err := render.Render(w, "home.html", pv); err != nil {
		b.log.Println(err)
		return
	}
}
