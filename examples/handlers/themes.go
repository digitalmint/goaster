package handlers

import (
	"net/http"

	"github.com/digitalmint/goaster/examples/pages"

	"github.com/a-h/templ"
	"github.com/digitalmint/goaster"
)

func HandleThemes(w http.ResponseWriter, r *http.Request) {
	toaster := goaster.NewToasterBuilder().WithAutoDismiss(false).Build()
	templ.Handler(pages.ThemesPage(toaster)).ServeHTTP(w, r)
}
