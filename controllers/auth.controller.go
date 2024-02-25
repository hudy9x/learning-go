package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		log.Fatal("error signing string token jwt")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"token":   tokenString,
	})
}

func ParseJWT(c *gin.Context) {
	token := c.GetHeader("authorization")

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    token,
	})
}

func ValidateUser(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
