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

	UserRoute := router.Group("/api/user")
	UserRoute.GET("/", controllers.AllUsers)
	UserRoute.POST("/", controllers.AddUser)
	UserRoute.POST("/update", controllers.UpdateUser)
	//TODO: make route protected
	PostRoute := router.Group("/api/post")

	PostRoute.GET("/", controllers.AllPosts)
	PostRoute.POST("/", controllers.AddPost)
	PostRoute.POST("/musicurl/update/:id", controllers.AddMusic)

	router.Run(":8080")
}
func main() {
	services.ConnectDatabase()
	RunRouter()
}
