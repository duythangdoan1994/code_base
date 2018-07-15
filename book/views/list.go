package bookviews

import "github.com/gin-gonic/gin"

func ListBook(c *gin.Context) {
	//var books []Book
	//err := GetAllBook(db, &books)
	//if err != nil {
	//	respondError(w, http.StatusInternalServerError, err.Error())
	//	return
	//}
	//respondJSON(w, http.StatusOK, books)
	c.JSON(200, gin.H{"ok": "Welcome to Chicago!"})
}