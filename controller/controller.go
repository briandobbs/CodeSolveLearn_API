package controller

import (
	"CodeSolveLearn_API/db"
	"CodeSolveLearn_API/service"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func Handler() *gin.Engine {

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
	articleController := NewArticleController(articleService)
	authorController := NewAuthorController(authorService)

	// Initialize the router
	router := gin.New()
	router.Use(gin.Recovery())

	v1 := router.Group("/v1")
	v1.Use()
	{
		// Define routes
		articles := v1.Group("/articles")
		{
			articles.GET("/", articleController.GetArticles)
			articles.GET("/:id", articleController.GetArticle)
			articles.POST("/", articleController.CreateArticle)
			//articles.PUT("/:id", articleController.UpdateArticle)
			//articles.DELETE("/:id", articleController.DeleteArticle)
		}

		authors := v1.Group("/authors")
		{
			authors.GET("/", authorController.GetAuthors)
			//authors.GET("/:id", authorController.GetAuthor)
			//authors.POST("/", authorController.CreateAuthor)
			//authors.PUT("/:id", authorController.UpdateAuthor)
			//authors.DELETE("/:id", authorController.DeleteAuthor)
		}
	}
	return router
}
