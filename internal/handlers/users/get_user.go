package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}
