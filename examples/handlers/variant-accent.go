package handlers

import (
	"net/http"

	"github.com/digitalmint/goaster/examples/pages"

	"github.com/a-h/templ"
	"github.com/digitalmint/goaster"
)

func HandleVariantAccent(w http.ResponseWriter, r *http.Request) {
	toaster := goaster.NewToaster(
		goaster.WithVariant(goaster.Accent),
		goaster.WithAutoDismiss(false),
	)
	templ.Handler(pages.VariantAccentLightPage(toaster)).ServeHTTP(w, r)
}
