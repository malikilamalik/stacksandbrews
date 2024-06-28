package server

import (
	"net/http"

	menuRoute "github.com/malikilamalik/stacksandbrews/services/catalog/internal/delivery/http/v1/menu"
)

func router(mux *http.ServeMux) {
	menuRoute := menuRoute.Init()
	mux.Handle("/menu/", menuRoute)
}
