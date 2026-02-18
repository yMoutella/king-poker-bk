package main

import (
	"github.com/ymoutella/king-poker-bk/internal/handler/users"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/users", users.GetUsers)
	router.GET("/users/:id", users.GetUser)
	router.POST("/users", users.CreateUser)
	router.Run()
}
