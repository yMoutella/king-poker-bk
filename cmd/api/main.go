package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ymoutella/king-poker-bk/internal/handler/users"
)

func main() {

	router := gin.Default()
	router.GET("/users", users.GetUsers)
	router.GET("/users/:id", users.GetUser)
	router.POST("/users", users.CreateUser)
	router.Run()
}
