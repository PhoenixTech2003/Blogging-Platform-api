package router

import (
	"net/http"

	"github.com/PhoenixTech2003/Blogging-Platform-api/internal/handlers"
)



func ArticleRouter () *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /", handlers.ApiConfig.CreateArticle)

	return mux
}