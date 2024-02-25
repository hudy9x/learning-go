package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	// .Bind function return an error
	// in this case: c.Bind(&body) means
	// if body is not JSON type then return error
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read the body",
		})

		return
	}

	// has the password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash the password",
		})
	}

	// create user

	c.JSON(http.StatusOK, gin.H{
		"success": "sign up successfully",
		"message": hash,
	})

}
