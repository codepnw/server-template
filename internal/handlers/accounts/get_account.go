package accounts

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAccount(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}
