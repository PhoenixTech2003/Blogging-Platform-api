package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type createArticleRequestBody struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type createArticleResponseBody struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
}

// CreateArticle godoc
// @Summary Create a new article
// @Description Create a new article with the provided title and content
// @Tags articles
// @Accept json
// @Produce json
// @Param article body createArticleRequestBody true "Article data"
// @Success 201 {object} createArticleResponseBody "Article created"
// @Router /articles [post]
func (cfg *ApiCfg) CreateArticle(w http.ResponseWriter, r *http.Request) {
	requestBody := createArticleRequestBody{}
	decoder := json.NewDecoder(r.Body)

	decoder.Decode(&requestBody)

	
}
