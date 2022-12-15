package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/abelblossom/todo/src/auth"
	"github.com/abelblossom/todo/src/db"
	"github.com/abelblossom/todo/src/models"
	"github.com/gin-gonic/gin"
)

func ListTodo(c *gin.Context) {
	claims := auth.GetClaims(c)

	var todos []models.Todo

	db.DB.Model(&models.Todo{UserID: claims.ID}).Find(&todos)

	c.JSON(200, gin.H{"data": todos})
}

func CreateToDo(c *gin.Context) {
	type createTodoType struct {
		Content string
	}
	claims := auth.GetClaims(c)

	body := createTodoType{}
	if err := c.ShouldBind(&body); err == nil && body.Content != "" {
		todo := models.Todo{Content: body.Content, UserID: claims.ID}
		db.DB.Create(&todo)
		c.JSON(http.StatusCreated, gin.H{"data": todo})
		c.Abort()
	} else {
		c.JSON(400, gin.H{"error": "Invalid Body"})
		c.Abort()
	}

}

func ToggleTodo(c *gin.Context) {
	claims := auth.GetClaims(c)

	if id, exist := c.Params.Get("id"); exist {
		todoId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			c.Abort()
			return
		}
		var todo models.Todo

		db.DB.First(&todo, todoId)

		if todo.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Todo Not Found"})
			c.Abort()
			return
		}

		if todo.UserID != claims.ID {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization failed"})
			c.Abort()
			return
		}

		fmt.Println(&todo)

		db.DB.Model(&todo).Update("completed", !todo.Completed)

		c.JSON(http.StatusCreated, gin.H{"data": &todo})
	}
}
