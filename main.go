package main

import (
	"zappin/controllers"
	"zappin/services"

	"github.com/gin-gonic/gin"
)

var PostRoute *gin.RouterGroup

func RunRouter() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	PostRoute := router.Group("/api/post")

	PostRoute.GET("/", controllers.AllPosts)

	router.Run(":8080")
}
func main() {
	services.ConnectDatabase()
	RunRouter()
}
