package users

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ymoutella/king-poker-bk/internal/database"
	"github.com/ymoutella/king-poker-bk/internal/domain"
)

func GetUsers(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "users handler",
	})
}

func GetUser(c *gin.Context) {
	var user GetUserParam
	if err := c.ShouldBindUri(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user id",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User returned",
		"data":    "data",
	})
}

func CreateUser(c *gin.Context) {
	var paramCreate CreateUserParam

	if err := c.ShouldBindJSON(&paramCreate); err != nil {

		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			errors := make(map[string]string)

			for _, fieldErr := range validationErrors {
				fieldName := strings.ToLower(fieldErr.Field())
				errors[fieldName] = fieldErr.Tag()
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid body request",
				"details": errors,
			})
			return

		}

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid body request",
			"error":   err.Error(),
		})
		return
	}

	db, err := database.PostgresDB()

	if err != nil {
		panic("Deu ruim na conexão")
	}

	db.AutoMigrate(domain.User{})

	var user domain.User = domain.User{
		Email:     paramCreate.Email,
		Password:  paramCreate.Password,
		FirstName: paramCreate.FirstName,
		LastName:  paramCreate.LastName,
	}

	result := db.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error in user creation",
			"error":   result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created",
		"data":    user,
	})
}
