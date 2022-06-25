package utils

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, status int, data interface{}, message string) {
	c.JSON(status, gin.H{
		"status":  status,
		"data":    data,
		"message": message,
	})
}
