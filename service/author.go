package service

import (
	"CodeSolveLearn_API/models"
	"gorm.io/gorm"
)

type AuthorService struct {
	db *gorm.DB
}

// NewAuthorService creates a new instance of AuthorService with the DB injected
func NewAuthorService(db *gorm.DB) *AuthorService {
	return &AuthorService{db: db}
}

// GetAllAuthors retrieves all authors from the database
func (s *AuthorService) GetAllAuthors() ([]models.Author, error) {
	var authors []models.Author
	result := s.db.Find(&authors)
	return authors, result.Error
}
