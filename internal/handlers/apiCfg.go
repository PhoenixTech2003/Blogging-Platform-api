package handlers

import (
	"database/sql"
	"log"
	"os"

	"github.com/PhoenixTech2003/Blogging-Platform-api/internal/database"
	"github.com/joho/godotenv"
)

type ApiCfg struct {
	db *database.Queries
}

var ApiConfig ApiCfg

func init() {

	_ = godotenv.Load()

	dbURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("an error occured while opening the database connection %v", err.Error())
		return
	}
	ApiConfig.db = database.New(db)

}
