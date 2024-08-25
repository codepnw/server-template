package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}
