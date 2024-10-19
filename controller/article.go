package controller

import (
	"CodeSolveLearn_API/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// GetAllArticles retrieves all articles from the database
func GetAllArticles(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var articles []models.Article

		// Query to get all articles along with the associated authors
		result := db.Preload("Author").Find(&articles)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, articles)
	}
}
