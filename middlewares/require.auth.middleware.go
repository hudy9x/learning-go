package middlewares

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(c *gin.Context) {
	log.Println("run in middleware")

	jwtToken := c.GetHeader("authorization")

	if jwtToken == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	log.Println("token is", jwtToken)

	token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		log.Println("parse jwt token error")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims := token.Claims.(jwt.MapClaims)

	log.Println(claims["email"], claims["photo"], claims["id"])

	c.Set("user", claims["id"])
	c.Next()
}
