package controller

import (
	"CodeSolveLearn_API/db"
	"CodeSolveLearn_API/internal/controller/appcontext"
	"github.com/gin-gonic/gin"
)

// Handler initializes the Gin router and passes the DB to all routes
func Handler(database db.Database) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	v1 := router.Group("/v1")
	v1.Use(func(c *gin.Context) {
		// Inject the database into the AppContext
		appCtx := &appcontext.AppContext{Context: c, DB: database}
		c.Set("appcontext", appCtx)
	})
	{
		v1.GET("/articles", appcontext.NewHandler(GetArticles))
	}
	return router
}
