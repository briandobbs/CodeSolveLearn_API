package main

import (
	"CodeSolveLearn_API/controller"
	"CodeSolveLearn_API/db"
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

	// Initialize the Gin router
	router := gin.Default()

	// Define the routes
	router.GET("/articles", controller.GetAllArticles(database))

	// Start the server
	router.Run(":9080")
}
