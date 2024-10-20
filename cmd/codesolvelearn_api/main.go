package main

import (
	"CodeSolveLearn_API/controller"
	"CodeSolveLearn_API/db"
	"CodeSolveLearn_API/service"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	// Initialize the database
	database, err := db.InitDB(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		3306,
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize services
	articleService := service.NewArticleService(database)
	authorService := service.NewAuthorService(database)

	// Initialize controllers
	articleController := controller.NewArticleController(articleService)
	authorController := controller.NewAuthorController(authorService)

	// Initialize the router
	router := gin.Default()

	// Define routes
	router.GET("/articles", articleController.GetArticles)
	router.GET("/authors", authorController.GetAuthors)

	// Start the server
	router.Run(":9080")
}
