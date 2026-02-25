package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/ymoutella/king-poker-bk/internal/auth"
	"github.com/ymoutella/king-poker-bk/internal/handler/login"
	"github.com/ymoutella/king-poker-bk/internal/handler/users"
)

func main() {

	router := gin.Default()
	router.POST("/login", login.Login)

	authorized := router.Group("/")
	authorized.Use(auth.AuthMiddleware())
	{
		authorized.GET("/users", users.GetUsers)
		authorized.GET("/users/:id", users.GetUser)
		authorized.POST("/users", users.CreateUser)
	}
	router.Run()
}
