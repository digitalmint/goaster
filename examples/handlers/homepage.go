package handlers

import (
	"net/http"

	"github.com/digitalmint/goaster/examples/pages"

	"github.com/a-h/templ"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	templ.Handler(pages.HomePage()).ServeHTTP(w, r)
}
