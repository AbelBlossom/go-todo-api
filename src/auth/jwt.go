package auth

import (
	"fmt"
	"time"

	"github.com/abelblossom/todo/src/db"
	"github.com/abelblossom/todo/src/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("super_secret key")

type JWTClaims struct {
	ID    uint
	Email string
	jwt.StandardClaims
}

type SlimUser struct {
	Email string `json:"email"`
	ID    uint   `json:"id"`
	Name  string `json:"name"`
}

func (c *JWTClaims) GetUser() SlimUser {

	user := SlimUser{}
	db.DB.Model(&models.User{}).Select("email", "id", "name").First(&user, c.ID)
	fmt.Println("user is", user)
	return user
}

func GenerateToken(user *models.User) (string, error) {
	claims := JWTClaims{}
	claims.ID = user.ID
	claims.Email = user.Email
	claims.ExpiresAt = time.Now().Add(60 * time.Minute).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if hash, err := token.SignedString(secretKey); err == nil {
		return hash, err
	} else {
		e := fmt.Errorf("could not sign %s", err.Error())
		return "", e
	}

}

func DecodeToken(token string) (*JWTClaims, error) {
	fmt.Println("token is", token)
	if t, err := jwt.ParseWithClaims(token, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	}); err == nil {
		claim := t.Claims.(*JWTClaims)

		if uint(claim.ExpiresAt) < uint(time.Now().Unix()) {
			return &JWTClaims{}, fmt.Errorf("%s", "Token Expired")
		}
		return claim, nil
	} else {
		return &JWTClaims{}, err
	}

}

func GetClaims(c *gin.Context) *JWTClaims {
	val, _ := c.Get("claims")
	return val.(*JWTClaims)
}
