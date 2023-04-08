package routes

import (
	userController "secure/challenge-2/infrastructure/restapi/controllers/user"
	"secure/challenge-2/infrastructure/restapi/middlewares"

	"github.com/gin-gonic/gin"
)

// UserRoutes is a function that contains all routes of the user
func UserRoutes(router *gin.RouterGroup, controller *userController.Controller) {
	routerAuth := router.Group("/user")

	// user routes without middleware
	{
		routerAuth.POST("", controller.NewUser)
	}

	// user routes using middleware
	routerAuth.Use(middlewares.AuthJWTMiddleware())
	{
		routerAuth.GET("/:id", controller.GetUsersByID)
		routerAuth.GET("", controller.GetAllUsers)
		routerAuth.PUT("/:id", controller.UpdateUser)
		routerAuth.DELETE("/:id", controller.DeleteUser)
	}
}
