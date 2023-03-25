package main

import (
	"secure/challenge-2/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init Router
	router := gin.Default()

	// Books Router
	routers.BookRouters(router)

	// Router Port
	err := router.Run(":4000")

	// Handle Router
	if err != nil {
		panic("Error When Running")
	}
}
