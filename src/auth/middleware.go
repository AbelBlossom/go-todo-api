package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare(c *gin.Context) {
	token := c.GetHeader("authorization")
	if token != "" {
		fmt.Println("Token Found")
		key := token[7:]
		if claim, err := DecodeToken(key); err == nil {
			c.Set("claims", claim)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "authorization failed",
			})
			c.Abort()
			return
		}
	}
}
