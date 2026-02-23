package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	usecases "github.com/ymoutella/king-poker-bk/internal/usecases/user"
)

func GetUsers(c *gin.Context) {
	var params GetAllUsersParam
	if err := c.BindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid query params",
			"error":   err.Error(),
		})
		return
	}

	us := usecases.FactoryGetAllUsersUS()
	fmt.Printf("Page: %d, PageSize: %d", params.Page, params.PageSize)
	users, err := us.Execute(params.Page, params.PageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching users",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Fetched successfully",
		"data":    users,
	})
}
