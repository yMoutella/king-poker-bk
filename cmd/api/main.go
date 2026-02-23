package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ymoutella/king-poker-bk/internal/auth"
	"github.com/ymoutella/king-poker-bk/internal/handler/users"
)

func main() {

	router := gin.Default()
	router.GET("/users", users.GetUsers)
	router.GET("/users/:id", users.GetUser)
	router.POST("/users", users.CreateUser)

	authorized := router.Group("/test")
	authorized.Use(auth.AuthMiddleware())
	{
		authorized.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "test",
			})
		})
	}
	router.Run()
}
