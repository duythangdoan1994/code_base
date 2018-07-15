package book

import (
	"github.com/gin-gonic/gin"
	"github.com/luciandd/core/book/views"
)

func RegisterBook(router *gin.RouterGroup) {
	router.GET("/", bookviews.ListBook)
}