package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMovie(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
