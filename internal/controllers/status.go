package controllers

import "github.com/gin-gonic/gin"

// Ping endpoint checks is service is available
func Ping(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "pong",
	})
}