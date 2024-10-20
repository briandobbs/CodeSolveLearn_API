package controller

import (
	"CodeSolveLearn_API/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ArticleController wraps the service for handling requests
type ArticleController struct {
	articleService *service.ArticleService
}

// NewArticleController creates a new instance of ArticleController
func NewArticleController(articleService *service.ArticleService) *ArticleController {
	return &ArticleController{articleService: articleService}
}

// GetArticles handles the GET request to fetch all articles
func (ac *ArticleController) GetArticles(c *gin.Context) {
	articles, err := ac.articleService.GetAllArticles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, articles)
}
