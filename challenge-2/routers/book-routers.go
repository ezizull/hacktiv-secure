package routers

import (
	"secure/challenge-2/controllers"

	"github.com/gin-gonic/gin"
)

func BookRouters(router *gin.Engine) {
	RBook := router.Group("/books")
	{
		RBook.GET("", controllers.GetBooks)
		RBook.GET("/:id", controllers.GetBook)
		RBook.POST("", controllers.AddBook)
		RBook.PUT("/:id", controllers.UpdateBook)
		RBook.DELETE("/:id", controllers.DeleteBook)
	}
}
