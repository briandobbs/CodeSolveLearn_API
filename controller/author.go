package controller

import (
	"CodeSolveLearn_API/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthorController wraps the service for handling requests
type AuthorController struct {
	authorService *service.AuthorService
}

// NewAuthorController creates a new instance of AuthorController
func NewAuthorController(authorService *service.AuthorService) *AuthorController {
	return &AuthorController{authorService: authorService}
}

// GetAuthors handles the GET request to fetch all authors
func (ac *AuthorController) GetAuthors(c *gin.Context) {
	authors, err := ac.authorService.GetAllAuthors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, authors)
}
