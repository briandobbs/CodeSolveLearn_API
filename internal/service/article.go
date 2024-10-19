package service

import (
	"CodeSolveLearn_API/db" // Import the Database interface
	"CodeSolveLearn_API/internal/entity"
	"context"
)

// ArticleService struct with dependency on Database interface
type ArticleService struct {
	db db.Database // Use the Database interface from the db package
}

// NewArticleService initializes ArticleService with a Database
func NewArticleService(database db.Database) *ArticleService {
	return &ArticleService{db: database}
}

// GetArticles retrieves all articles from the database
func (s *ArticleService) GetArticles(ctx context.Context) ([]entity.Article, error) {
	query := "SELECT id, title, description, author, body FROM Article"
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []entity.Article
	for rows.Next() {
		var a entity.Article
		if err := rows.Scan(&a.ID, &a.Title, &a.Description, &a.Author, &a.Body); err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}
	return articles, rows.Err()
}
