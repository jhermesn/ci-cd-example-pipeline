package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Count(c *gin.Context) {
	text := c.Query("text")
	if text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "text query param is required"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": len(strings.Fields(text))})
}
