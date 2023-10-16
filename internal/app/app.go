package app

import (
	// "fmt"
	"io"
	"log"
	"net/http"

	"github.com/The-Gleb/url-shortener/internal/storage"
	"github.com/go-chi/chi/v5"
)

type app struct {
	storage storage.Repository
}

type App interface {
	GetShortenedURL(rw http.ResponseWriter, r *http.Request)
	GetFullURL(rw http.ResponseWriter, r *http.Request)
}

func NewApp(s storage.Repository) *app {
	return &app{
		storage: s,
	}
}

func (a *app) GetShortenedURL(rw http.ResponseWriter, r *http.Request) {
	url, err := io.ReadAll(r.Body)
	log.Printf("THE URL IS %s", url)
	defer r.Body.Close()
	if err != nil {
		http.Error(rw, "Couldn`t read url", http.StatusBadRequest)
	}
	id := ShortenURL(string(url))
	log.Printf("THE ID IS %s", id)
	a.storage.AddURL(id, string(url))
	rw.Header().Set("Content-Type", "text/plain")
	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte("http://localhost:8080/" + id))
}

func (a *app) GetFullURL(rw http.ResponseWriter, r *http.Request) {
	log.Println("Start handler")
	id := chi.URLParam(r, "id")
	log.Println("Call Get URL")
	url, err := a.storage.GetURL(id)
	if err != nil {
		http.Error(rw, "url not found", http.StatusBadRequest)
	}
	rw.Header().Add("Location", url)
	rw.WriteHeader(http.StatusTemporaryRedirect)
	log.Println("Finish handler")
}
func ShortenURL(url string) string {
	return "EwHXdJfB"
}
