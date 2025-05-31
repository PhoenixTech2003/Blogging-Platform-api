package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/swag" // swagger general

	// Import the generated docs
	"log"
	"net/http"

	_ "github.com/PhoenixTech2003/Blogging-Platform-api/docs"
	"github.com/PhoenixTech2003/Blogging-Platform-api/internal/router"
)

// @title Blogging Platform API
// @version 1.0
// @description A RESTful API for a blogging platform
// @host localhost:8081
// @BasePath /v1/api

func main() {
	godotenv.Load()
	baseURL := os.Getenv("BASE_URL")
	mux := http.NewServeMux()

	// Serve Swagger documentation
	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%v/swagger/doc.json", baseURL)), // The url pointing to API definition
	))

	mux.Handle("/v1/api/articles/", http.StripPrefix("/v1/api/articles", router.ArticleRouter()))

	addr := ":8081"
	server := http.Server{
		Handler:           mux,
		Addr:              addr,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("server is listening at http://localhost%v", addr)
	log.Printf("Swagger UI is available at %v/swagger/index.html", baseURL)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
