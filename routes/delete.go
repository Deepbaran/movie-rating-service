package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteMovie(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": "pong",
	})
}
