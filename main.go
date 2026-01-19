package main

import (
	"pharmaciano/controllers"
	"pharmaciano/initializers"
	"pharmaciano/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}
func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/signup", controllers.Signup)
	router.POST("/signin", controllers.Sigin)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)
	router.Run()
}
