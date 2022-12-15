package controllers

import (
	"fmt"
	"net/http"

	"github.com/abelblossom/todo/src/auth"
	"github.com/abelblossom/todo/src/db"
	"github.com/abelblossom/todo/src/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	user := models.User{}

	if err := c.ShouldBind(&user); err == nil {
		db.DB.Create(&user)
	} else {
		panic(err)
	}

	c.JSON(200, gin.H{"user": user})
}

func Login(c *gin.Context) {
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

}
