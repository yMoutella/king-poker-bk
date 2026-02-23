package users

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ymoutella/king-poker-bk/internal/domain"
	usecases "github.com/ymoutella/king-poker-bk/internal/usecases/user"
)

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

	uc, err := usecases.FactoryCreateUserUS()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error in user creation",
			"error":   err.Error,
		})
		return
	}

	user := &domain.User{
		Email:     paramCreate.Email,
		Password:  paramCreate.Password,
		FirstName: paramCreate.FirstName,
		LastName:  paramCreate.LastName,
	}

	createdUser, err := uc.Execute(user)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Error in user creation",
			"error":   err.Error,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created",
		"data":    createdUser,
	})
}
