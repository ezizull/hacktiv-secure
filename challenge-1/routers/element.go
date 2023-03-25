package routers

import (
	"secure/challenge-1/controllers"

	"github.com/gin-gonic/gin"
)

func ElementRouters(router *gin.Engine) {
	RBook := router.Group("/posts")
	{
		RBook.POST("", controllers.PostElement)
	}
}