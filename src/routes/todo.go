package routes

import (
	"github.com/abelblossom/todo/src/auth"
	"github.com/abelblossom/todo/src/controllers"
	gin "github.com/gin-gonic/gin"
)

func MakeTodoRoute(r *gin.RouterGroup) {
	r.Use(auth.AuthMiddleWare)
	r.GET("/", controllers.ListTodo)
	r.POST("/", controllers.CreateToDo)
	r.GET("/:id", controllers.GetTodo)
	r.PUT("/:id/toggle", controllers.ToggleTodo)
	r.PUT("/:id", controllers.EditTodo)
	r.DELETE("/:id", controllers.DeleteTodo)
}
