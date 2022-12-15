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

	err := c.ShouldBind(&user)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid Content"})
		return
	}
	var userCount int64
	db.DB.First(&models.User{}, "email = ?", user.Email).Count(&userCount)
	if userCount > 0 {
		c.AbortWithStatusJSON(400, gin.H{"error": "User Already Exists"})
		return
	}
	db.DB.Create(&user)

	c.AbortWithStatusJSON(200, gin.H{"user": user})
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
					c.AbortWithStatusJSON(200, gin.H{"token": token})
					return
				} else {
					fmt.Println(err)
				}
			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Wrong Password"})
				return
			}
		}
	}
	c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"error": "failed"})

}
