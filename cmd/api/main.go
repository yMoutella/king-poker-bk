package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ymoutella/king-poker-bk/internal/handler/users"
)

func main() {

	// db, err := database.PostgresDB()

	// if err != nil {
	// 	panic("Deu ruim na conexão")
	// }

	// db.AutoMigrate(domain.User{})

	router := gin.Default()
	router.GET("/users", users.GetUsers)
	router.GET("/users/:id", users.GetUser)
	router.POST("/users", users.CreateUser)
	router.Run()
}
