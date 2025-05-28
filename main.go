package main

import (
	"encoding/json"
	"github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger" // swagger embed files
	_ "github.com/swaggo/swag"        // swagger general
	// Import the generated docs
	_ "github.com/PhoenixTech2003/Blogging-Platform-api/docs"
	"log"
	"net/http"
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
		Message: "Hello world",
	}

	dat, _ := json.Marshal(response)
	w.Write(dat)
}

func main() {
	mux := http.NewServeMux()
	
	// Serve Swagger documentation
	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8081/swagger/doc.json"), // The url pointing to API definition
	))

	// Register handlers
	mux.HandleFunc("/", homeHandler)

	addr := ":8081"
	server := http.Server{
		Handler: mux,
		Addr:    addr,
	}
	log.Printf("server is listening at http://localhost%v", addr)
	log.Printf("Swagger UI is available at http://localhost%v/swagger/index.html", addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
