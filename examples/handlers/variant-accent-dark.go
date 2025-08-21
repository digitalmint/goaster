package handlers

import (
	"net/http"

	"github.com/digitalmint/goaster/examples/pages"

	"github.com/a-h/templ"
	"github.com/digitalmint/goaster"
)

func HandleVariantAccentDark(w http.ResponseWriter, r *http.Request) {
	toaster := goaster.NewToaster(
		goaster.WithVariant(goaster.AccentDark),
		goaster.WithAutoDismiss(false),
	)
	templ.Handler(pages.VariantAccentLightPage(toaster)).ServeHTTP(w, r)
}
