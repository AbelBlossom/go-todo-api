package routes

import (
	"github.com/abelblossom/todo/src/auth"
	gin "github.com/gin-gonic/gin"
)

func MakeTodoRoute(r *gin.RouterGroup) {
	r.Use(auth.AuthMiddleWare)
	r.GET("/", func(c *gin.Context) {

		val, exists := c.Get("claims")
		if exists {
			v := val.(*auth.JWTClaims)
			// id := v.Email
			user := v.GetUser()
			c.JSON(200, gin.H{"data": v, "user": user})
			c.Abort()
			return
		}
		c.JSON(500, gin.H{"data": "no claims"})
		c.Next()
	})
}
