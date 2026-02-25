package login

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	usecases "github.com/ymoutella/king-poker-bk/internal/usecases/auth"
)

func Login(c *gin.Context) {

	var params AuthParams

	if err := c.ShouldBindJSON(&params); err != nil {
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
	}

	fmt.Print(params.Email, params.Password)
	us := usecases.FactoryAuthUseCase()

	token, err := us.Execute(params.Email, params.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success!",
		"token":   token,
	})

}
