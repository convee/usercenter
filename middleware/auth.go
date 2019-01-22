package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := c.Cookie("uid")
		if err != nil || uid == "" {
			c.Redirect(http.StatusMovedPermanently, "/member/login")
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("uid", uid)
	}

}