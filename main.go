package main

import (
	"fmt"
	"net/http"
	_ "reflect"

	"github.com/abelblossom/todo/src/controllers"
	"github.com/abelblossom/todo/src/db"
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

	server.POST("/user", controllers.CreateUser)

	server.POST("/login", controllers.Login)

	err := server.Run()
	if err != nil {
		fmt.Println("server started")
	}

}
