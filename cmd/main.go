package main

import (
	routes "event-server/api/route"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	routes.SetupRoutes(router)

	router.Run(":8080")
}
