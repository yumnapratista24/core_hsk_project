package handler

import (
	"core_hsk_project/errors"

	"github.com/gin-gonic/gin"
)

// handleError is a helper function to handle errors consistently
func HandleError(c *gin.Context, err errors.CustomError) {
	c.JSON(err.Code, gin.H{
		"success": false,
		"error": gin.H{
			"code":    err.Code,
			"message": err.Message,
		},
	})
}
