package main

import (
	"github.com/The-Gleb/url-shortener/internal/app"
	"github.com/The-Gleb/url-shortener/internal/server"
	"github.com/The-Gleb/url-shortener/internal/storage"
)

func main() {
	storage := storage.New()
	app := app.NewApp(storage)
	s := server.NewServer(":8080", app)

	server.RunServer(s)
}
