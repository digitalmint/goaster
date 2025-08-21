package handlers

import (
	"net/http"

	"github.com/digitalmint/goaster/examples/pages"

	"github.com/a-h/templ"
	"github.com/digitalmint/goaster"
)

func HandleSingle(w http.ResponseWriter, r *http.Request) {
	result := true
	toaster := goaster.ToasterDefaults()
	templ.Handler(pages.SingleToastPage(result, toaster)).ServeHTTP(w, r)
}
