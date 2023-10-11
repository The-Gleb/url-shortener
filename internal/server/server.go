package server

import (
	"net/http"

	"github.com/The-Gleb/url-shortener/internal/app"
	"github.com/go-chi/chi/v5"
)

func NewServer(address string, app app.App) *http.Server {
	router := chi.NewRouter()
	SetUpRoutes(router, app)
	return &http.Server{
		Addr:    address,
		Handler: router,
	}
}
func SetUpRoutes(router *chi.Mux, app app.App) {
	router.Post("/", app.GetShortenedURL)
	router.Get("/{id}", app.GetFullURL)
}

func RunServer(server *http.Server) error {
	return server.ListenAndServe()
}
