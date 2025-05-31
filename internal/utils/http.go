package utils

import (
	"log"
	"net/http"
)

func RespondWithJson (w http.ResponseWriter, data []byte, statusCode int)  {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(statusCode)
	_, err := w.Write(data)
	if err != nil {
		log.Printf("an error occured while writing response: %v", err.Error())
		return 
	}
}