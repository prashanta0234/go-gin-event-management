package routes

import (
	"event-server/internal"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	// v1 := router.Group("/api/v1")

	router.GET("/", func(c *gin.Context) {
		// c.JSON(200, gin.H{"message": "Hello World"})
		internal.SendSuccess(c, "Hello World", nil)
	})
}
