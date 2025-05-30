package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/swaggo/http-swagger"
	_ "github.com/swaggo/swag" // swagger general

	// Import the generated docs
	"log"
	"net/http"

	_ "github.com/PhoenixTech2003/Blogging-Platform-api/docs"
)

// @title Blogging Platform API
// @version 1.0
// @description A RESTful API for a blogging platform
// @host localhost:8081
// @BasePath /

// @Summary Home endpoint
// @Description Returns a hello world message
// @Produce json
// @Success 200 {object} map[string]string
// @Router / [get]
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := struct {
		Message string `json:"message"`
	}{
		Message: "Hello chichi",
	}

	dat, _ := json.Marshal(response)
	_, err := w.Write(dat)
	if err != nil {
		log.Printf("an error occured while writing data %v", err.Error())
	}
}

func main() {
	godotenv.Load()
	baseURL := os.Getenv("BASE_URL")
	mux := http.NewServeMux()

	// Serve Swagger documentation
	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%v/swagger/doc.json",baseURL)), // The url pointing to API definition
	))

	// Register handlers
	mux.HandleFunc("/", homeHandler)

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
