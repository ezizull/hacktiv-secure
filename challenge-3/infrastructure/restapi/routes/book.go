// Package routes contains all routes of the application
package routes

import (
	bookController "secure/challenge-3/infrastructure/restapi/controllers/book"
	"secure/challenge-3/infrastructure/restapi/middlewares"

	"github.com/gin-gonic/gin"
)

// UserRoutes is a function that contains all routes of the book
func BookRoutes(router *gin.RouterGroup, controller *bookController.Controller) {
	routerBook := router.Group("/books")

	// authentication
	routerBook.Use(middlewares.AuthJWTMiddleware())
	{
		routerBook.GET("", controller.GetAllBooks)
		routerBook.GET("/:id", controller.GetBookByID)
		routerBook.POST("", controller.NewBook)

	}

	// admin role
	routerBook.Use(middlewares.AuthRoleMiddleware([]string{"admin"}))
	{
		routerBook.PUT("/:id", controller.UpdateBook)
		routerBook.DELETE("/:id", controller.DeleteBook)
	}

}
