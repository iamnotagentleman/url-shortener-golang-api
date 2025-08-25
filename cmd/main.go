package main

import (
	"awesome-url-shortener/internal/config"
	"awesome-url-shortener/internal/handlers"
	"awesome-url-shortener/internal/models"
	"awesome-url-shortener/internal/service"
	"awesome-url-shortener/internal/store/keyval"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadEnvVars()

	if err != nil {
		log.Fatal(err)
	}
	keyValStore := keyval.NewKeyValueStore(cfg)
	service := service.NewService(keyValStore, cfg.Common)
	handler := handlers.NewHandler(service)

	mux := http.NewServeMux()

	mux.Handle("POST /url/", handlers.WithHTTPIn(handler.ShortURl, models.UrlShortCreateInput{}))
	mux.HandleFunc("GET /go/{key}", handler.ResolveShortUrl)
	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
