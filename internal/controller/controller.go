package controller

import "github.com/gin-gonic/gin"

func Handler() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
}
