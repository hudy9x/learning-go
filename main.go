package main

import (
	"example/go-web/controllers"
	"example/go-web/initializers"
	"example/go-web/middlewares"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// this function will be loaded automatically at startup
func init() {
	log.Println("Initialize the application !!!!!!!!")
	initializers.LoadEnvVars()
	_, err := initializers.ConnectDatabase()

	// log.Println(client)

	if err != nil {
		log.Fatal("error connection db")
	}

	// initializers.ScanApp()
	initializers.GetIpNPcName()
	// initializers.InsertAMovie(client)
}

func main() {

	fmt.Println("somthing else")

	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong 2",
		})
	})

	r.GET("/sign-up", controllers.Signup)
	r.GET("/jwt", controllers.GenerateJWT)
	r.GET("/validate", middlewares.RequireAuth, controllers.ValidateUser)

	// r.Run()
}
