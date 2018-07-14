package api

import (
	"github.com/gin-gonic/gin"
	"github.com/luciandd/core/book"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// register version api
	v1 := router.Group("api/v1")

	// Register app api
	book.RegisterBook(v1.Group("/book"))

	return router
}