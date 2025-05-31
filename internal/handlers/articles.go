package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/PhoenixTech2003/Blogging-Platform-api/internal/database"
	"github.com/PhoenixTech2003/Blogging-Platform-api/internal/utils"
)

type createArticleRequestBody struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type createArticleResponseBody struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type createArticleResponseError struct {
	Message string `json:"message"`
}

// CreateArticle godoc
// @Summary Create a new article
// @Description Create a new article with the provided title and content
// @Tags Articles
// @Accept json
// @Produce json
// @Param article body createArticleRequestBody true "Article data"
// @Success 201 {object} createArticleResponseBody "Article created"
// @Failure 500 {object} createArticleResponseError "Internal Server Error"
// @Router /articles/ [post]
func (cfg *ApiCfg) CreateArticle(w http.ResponseWriter, r *http.Request) {
	requestBody := createArticleRequestBody{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&requestBody)
	if err != nil {
		log.Printf("an error occured while decoding the request: %v", err.Error())
		errorResponse := createArticleResponseError{
			Message: "an error occured while creating the post",
		}
		dat, err := json.Marshal(errorResponse)
		if err != nil {
			log.Fatal("an error occured while marshalling json")
			return
		}

		utils.RespondWithJson(w, dat, http.StatusInternalServerError)

	}
	createArticleParams := database.CreateArticleParams{
		Title:   requestBody.Title,
		Content: requestBody.Content,
	}
	articleData, err := cfg.db.CreateArticle(r.Context(), createArticleParams)
	if err != nil {
		log.Printf("an error occured while creating the post: %v", err.Error())
		errorResponse := createArticleResponseError{
			Message: "an error occured while creating the post",
		}
		dat, err := json.Marshal(errorResponse)
		if err != nil {
			log.Fatal("failed to create article")
			return
		}

		utils.RespondWithJson(w, dat, http.StatusInternalServerError)

	}

	responseBody := createArticleResponseBody{
		ID:        articleData.ID.String(),
		Title:     articleData.Title,
		Content:   articleData.Content,
		CreatedAt: articleData.CreatedAt.Time,
		UpdatedAt: articleData.UpdatedAt.Time,
	}

	dat, err := json.Marshal(responseBody)
	if err != nil {
		log.Printf("an error occured while marshalling the post: %v", err.Error())
		errorResponse := createArticleResponseError{
			Message: "an error occured while creating the post",
		}
		dat, err := json.Marshal(errorResponse)
		if err != nil {
			log.Fatal("failed to marshal article")
			return
		}

		utils.RespondWithJson(w, dat, http.StatusInternalServerError)

	}

	utils.RespondWithJson(w, dat, http.StatusOK)

}
