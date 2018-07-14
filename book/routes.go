package book

import "github.com/gin-gonic/gin"

func RegisterBook(router *gin.RouterGroup) {
	router.GET("/", ListBook)
}