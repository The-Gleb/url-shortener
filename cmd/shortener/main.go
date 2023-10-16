package main

import (
	"github.com/The-Gleb/url-shortener/internal/app"
	"github.com/The-Gleb/url-shortener/internal/config"
	"github.com/The-Gleb/url-shortener/internal/server"
	"github.com/The-Gleb/url-shortener/internal/storage"
)

//chi is already in use

func main() {
	config := config.NewConfigFromFlags()
	storage := storage.New()
	app := app.NewApp(storage, config.BaseAddress)
	s := server.NewServer(config.ServerAddress, app)

	server.RunServer(s)
}
