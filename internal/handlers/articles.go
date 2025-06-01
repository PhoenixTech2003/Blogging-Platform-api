package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PhoenixTech2003/Blogging-Platform-api/internal/database"
	"github.com/google/uuid"

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

type responseError struct {
	Message string `json:"message"`
}

type article struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateArticle godoc
// @Summary Create a new article
// @Description Create a new article with the provided title and content
// @Tags Articles
// @Accept json
// @Produce json
// @Param article body createArticleRequestBody true "Article data"
// @Success 201 {object} createArticleResponseBody "Article created"
// @Failure 500 {object} responseError "Internal Server Error"
// @Router /articles/ [post]
func (cfg *ApiCfg) CreateArticle(w http.ResponseWriter, r *http.Request) {
	requestBody := createArticleRequestBody{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&requestBody)
	if err != nil {
		log.Printf("an error occured while decoding the request: %v", err.Error())
		errorResponse := responseError{
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
		errorResponse := responseError{
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
		errorResponse := responseError{
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

// GetArticles godoc
// @Summary Gets all articles
// @Description Fetches all blog articles or blog articles that  match the query field
// @Tags Articles
// @Param query query string false "optional search parameter, it is case insensitive"
// @Accept json
// @Produce json
// @Success 200 {array} article  "All articles"
// @Failure 500 {object} responseError "Internal Server Error"
// @Router /articles/ [get]
func (cfg *ApiCfg) GetArticles(w http.ResponseWriter, r *http.Request) {
	queryParameter := r.URL.Query().Get("query")
	articles := make([]article, 0)
	articlesData, err := cfg.db.GetArticles(r.Context(), fmt.Sprintf("%%%v%%", queryParameter))
	if err != nil {
		log.Printf("an error occured while fetching all articles: %v", err.Error())
		errorResponse := responseError{
			Message: "an error occured while fetching all articles",
		}
		errorData, _ := json.Marshal(errorResponse)
		utils.RespondWithJson(w, errorData, http.StatusInternalServerError)
		return
	}

	for _, articleData := range articlesData {

		articles = append(articles, article{
			ID:        articleData.ID.String(),
			Title:     articleData.Title,
			Content:   articleData.Content,
			CreatedAt: articleData.CreatedAt.Time,
			UpdatedAt: articleData.UpdatedAt.Time,
		})

	}

	dat, err := json.Marshal(articles)
	if err != nil {
		log.Printf("an error occured while marshalling all articles: %v", err.Error())
		errorResponse := responseError{
			Message: "an error occured while fetching all articles",
		}
		errorData, _ := json.Marshal(errorResponse)
		utils.RespondWithJson(w, errorData, http.StatusInternalServerError)
		return
	}

	utils.RespondWithJson(w, dat, http.StatusOK)

}

// GetArticleById godoc
// @Summary fetches article  by ID
// @Tags Articles
// @Description Fetches an article that matches the id specified in the url
// @Param articleId path string true "The id of the article"
// @Success 200 {object} article "fetched article"
// @Failure 500 {object} responseError "Internal server error"
// @Failure 404 {object} responseError "Article not found"
// @Router /articles/{articleId} [get]
func (cfg *ApiCfg) GetArticleById(w http.ResponseWriter, r *http.Request){
	articleId := r.PathValue("articleId")
	parsedUUID, err :=uuid.Parse(articleId)
	if err != nil {
		log.Printf("an error occured while parsing uuid of  article of Id %v: %v",articleId, err.Error())
		errorResponse := responseError{
			Message: fmt.Sprintf("an error while fetching artlce of id %v", articleId),
		}
		errorData, _ := json.Marshal(errorResponse)
		utils.RespondWithJson(w, errorData, http.StatusInternalServerError)
		return
	}
	dat, err := cfg.db.GetArtidleByID(r.Context(),parsedUUID)
	if err != nil {
		if err.Error() == "sql: no rows in result set"{
		log.Printf("an error occured while fetching  article: %v", err.Error())
		errorResponse := responseError{
			Message: "article not found",
		}
		errorData, _ := json.Marshal(errorResponse)
		utils.RespondWithJson(w, errorData, http.StatusNotFound)
		return	
		}
		log.Printf("an error occured while fetching the  article of Id %v: %v",articleId, err.Error())
		errorResponse := responseError{
			Message: fmt.Sprintf("an error while fetching artlce of id %v", articleId),
		}
		errorData, _ := json.Marshal(errorResponse)
		utils.RespondWithJson(w, errorData, http.StatusInternalServerError)
		return
	}

	responseBody :=article{
		ID: dat.ID.String(),
		Title: dat.Title,
		Content: dat.Content,
		CreatedAt: dat.CreatedAt.Time,
		UpdatedAt: dat.UpdatedAt.Time,
	}

	responseData, err := json.Marshal(responseBody)
	if err != nil {
		
		log.Printf("an error occured while marshalling  article: %v", err.Error())
		errorResponse := responseError{
			Message: "an error occured while fetching article",
		}
		errorData, _ := json.Marshal(errorResponse)
		utils.RespondWithJson(w, errorData, http.StatusInternalServerError)
		return
	}

	utils.RespondWithJson(w,responseData,http.StatusOK)
	
}