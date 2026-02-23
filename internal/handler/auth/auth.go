package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Auth(c *gin.Context) {

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

	// Call the auth use case (have to write)

	return
}
