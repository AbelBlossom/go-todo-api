package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/abelblossom/todo/src/auth"
	"github.com/abelblossom/todo/src/db"
	"github.com/abelblossom/todo/src/models"
	"github.com/gin-gonic/gin"
)

func ListTodo(c *gin.Context) {
	claims := auth.GetClaims(c)

	var todos []models.Todo

	db.DB.Model(&models.Todo{}).Find(&todos, "user_id == ?", claims.ID)

	c.AbortWithStatusJSON(200, gin.H{"data": todos})
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
		c.AbortWithStatusJSON(http.StatusCreated, gin.H{"data": todo})
	} else {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid Body"})
	}

}

func GetTodo(c *gin.Context) {
	claims := auth.GetClaims(c)
	id, exist := c.Params.Get("id")
	if _, err := strconv.Atoi(id); !exist && err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	todo := models.Todo{}

	db.DB.First(&todo, id)
	if todo.ID == 0 {
		c.AbortWithStatusJSON(400, gin.H{"error": "Todo Not Found"})
		return
	}
	if todo.UserID != claims.ID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization Failed"})
		return
	}
	c.AbortWithStatusJSON(200, gin.H{"data": todo})
}

func ToggleTodo(c *gin.Context) {
	claims := auth.GetClaims(c)

	if id, exist := c.Params.Get("id"); exist {
		todoId, err := strconv.Atoi(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}
		var todo models.Todo

		db.DB.First(&todo, todoId)

		if todo.ID == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Todo Not Found"})
			return
		}

		if todo.UserID != claims.ID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization failed"})
			return
		}

		fmt.Println(&todo)

		db.DB.Model(&todo).Update("completed", !todo.Completed)

		c.AbortWithStatusJSON(http.StatusCreated, gin.H{"data": &todo})
		return
	}

	c.AbortWithStatusJSON(400, gin.H{"error": "Error"})
}

func EditTodo(c *gin.Context) {
	claims := auth.GetClaims(c)
	id, exist := c.Params.Get("id")
	if !exist {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid ID"})

		return
	}
	type editType struct{ Content string }
	body := editType{}
	err := c.ShouldBind(&body)
	if err != nil || strings.TrimSpace(body.Content) == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid Content"})
		return
	}
	todo := models.Todo{}
	db.DB.Model(&todo).Find(&todo, id)
	if todo.ID == 0 {
		c.AbortWithStatusJSON(404, gin.H{"error": "Todo Not Found"})
		return
	}
	if todo.UserID != claims.ID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization failed"})
		return
	}

	db.DB.Model(&todo).Update("content", strings.TrimSpace(body.Content))

	c.AbortWithStatusJSON(200, gin.H{"data": todo})

}

func DeleteTodo(c *gin.Context) {
	claims := auth.GetClaims(c)
	id, exist := c.Params.Get("id")
	if !exist {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	db.DB.Where("user_id = ?", claims.ID).Delete(&models.Todo{}, id)

	c.AbortWithStatusJSON(200, gin.H{"success": "Deleted"})
}
