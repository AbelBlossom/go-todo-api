package main

import (
	"fmt"
	"net/http"
	_ "reflect"

	"github.com/abelblossom/todo/src/auth"
	"github.com/abelblossom/todo/src/db"
	"github.com/abelblossom/todo/src/models"
	"github.com/abelblossom/todo/src/routes"
	gin "github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	db.ConnectToDataBase()
	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello"})
	})

	todoGroup := server.Group("/todo")

	routes.MakeTodoRoute(todoGroup)

	server.POST("/user", func(c *gin.Context) {

		user := models.User{}

		if err := c.ShouldBind(&user); err == nil {
			db.DB.Create(&user)
		} else {
			panic(err)
		}

		c.JSON(200, gin.H{"user": user})
	})

	server.POST("/login", func(c *gin.Context) {
		type loginDet struct {
			Email    string
			Password string
		}
		details := loginDet{}
		if err := c.ShouldBind(&details); err == nil {
			user := models.User{Email: details.Email}
			db.DB.Where(user).First(&user)
			if user.ID != 0 {
				if user.ComparePassword(details.Password) {
					if token, err := auth.GenerateToken(&user); err == nil {
						c.JSON(200, gin.H{"token": token})
						c.Abort()
						return
					} else {
						fmt.Println(err)
					}
				} else {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Wrong Password"})
					c.Abort()
					return
				}
			}
		}
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed"})
		c.Abort()

	})

	err := server.Run()
	if err != nil {
		fmt.Println("server started")
	}

}
