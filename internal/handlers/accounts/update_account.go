package accounts

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateAccount(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}
