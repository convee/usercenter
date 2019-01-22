package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleMessage(message string, c *gin.Context)  {
	c.HTML(http.StatusNotFound, "error.html", gin.H{
		"message": message,
	})
}

func Handle404(c *gin.Context)  {
	HandleMessage("Page Not Found...", c)
}