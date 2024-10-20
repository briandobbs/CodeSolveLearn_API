package service

import (
	"CodeSolveLearn_API/models"
	"gorm.io/gorm"
)

type ArticleService struct {
	db *gorm.DB
}

// NewArticleService creates a new instance of ArticleService with the DB injected
func NewArticleService(db *gorm.DB) *ArticleService {
	return &ArticleService{db: db}
}

// GetAllArticles retrieves all articles from the database
func (s *ArticleService) GetAllArticles() ([]models.Article, error) {
	var articles []models.Article
	result := s.db.Preload("Author").Find(&articles)
	return articles, result.Error
}
