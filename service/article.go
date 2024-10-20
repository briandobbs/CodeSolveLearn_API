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
	//result := s.db.Preload("Author").Find(&articles)
	result := s.db.Find(&articles)
	return articles, result.Error
}

// GetArticleByID retrieves an article by its ID
func (s *ArticleService) GetArticleByID(id string) (models.Article, error) {
	var article models.Article
	//result := s.db.Preload("Author").First(&article, id)
	result := s.db.First(&article, id)
	return article, result.Error
}

// CreateArticle creates a new article in the database
func (s *ArticleService) CreateArticle(article models.Article) (models.Article, error) {
	result := s.db.Create(&article)
	return article, result.Error
}
