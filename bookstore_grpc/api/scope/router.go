package api

import (
	"github.com/gin-gonic/gin"
)

// GinRouter is router that API method router.
func GinRouter() *gin.Engine {
	router := gin.Default()

	// Simple group: v1
	book := router.Group("/book")
	{
		book.GET("/list", getBookList)
		book.GET("/search/:title", getBookByTitle)
		//		book.GET("/:id", getBookByID)
		book.PATCH("/:id", updateBookInfo)
		book.POST("/insert", addBookInfo)
		book.DELETE("/:id", delBookInfo)
	}

	return router
}
