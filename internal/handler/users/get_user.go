package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	usecases "github.com/ymoutella/king-poker-bk/internal/usecases/user"
)

func GetUser(c *gin.Context) {
	var params GetUserParam
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user id",
		})
		return
	}

	us := usecases.FactoryGetByIdUS()
	id, err := uuid.Parse(params.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user id format",
			"error":   err.Error(),
		})
		return
	}

	user, err := us.Execute(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching user",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User returned",
		"data":    user,
	})
}
