package service_test

import (
	"CodeSolveLearn_API/internal/service"
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Mock database that satisfies the db.Database interface
func TestGetArticles(t *testing.T) {
	// Create a mock database connection
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()

	// Setup the mock rows to return from the database
	rows := sqlmock.NewRows([]string{"id", "title", "description", "author", "body"}).
		AddRow(1, "Article 1", "Description 1", "Author 1", "Body 1")

	// Expect the query to be run
	mock.ExpectQuery("SELECT id, title, description, author, body FROM Article").
		WillReturnRows(rows)

	// Inject the mock database into the service
	articleService := service.NewArticleService(mockDB)

	// Execute the service method
	articles, err := articleService.GetArticles(context.Background())

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, articles, 1)
}
