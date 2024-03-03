package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rookie623/gin-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// listen and serve on 0.0.0.0:8080
	if err := router.Run(); err != nil {
		return
	}
}
