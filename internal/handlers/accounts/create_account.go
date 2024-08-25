package accounts

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"success": true})
}
