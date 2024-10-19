package controller

import (
	"CodeSolveLearn_API/internal/controller/appcontext"
	"CodeSolveLearn_API/internal/service"
	"context"
)

func GetArticles(c *appcontext.AppContext) {
	// Initialize the ArticleService with the database from the context
	articleService := service.NewArticleService(c.DB)

	// Call the service to get all articles and handle the response
	articles, err := articleService.GetArticles(context.Background())
	c.HandleResponse(articles, err)
}
