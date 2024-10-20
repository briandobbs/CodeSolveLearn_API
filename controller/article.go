package controller

import (
	"CodeSolveLearn_API/models"
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

func (ac *ArticleController) GetArticle(c *gin.Context) {
	id := c.Param("id")
	article, err := ac.articleService.GetArticleByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, article)
}

func (ac *ArticleController) CreateArticle(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdArticle, err := ac.articleService.CreateArticle(article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdArticle)
}
