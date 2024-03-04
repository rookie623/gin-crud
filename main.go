package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rookie623/gin-crud/controllers"
	"github.com/rookie623/gin-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.PostIndex)
	r.GET("/posts/:id", controllers.PostGet)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)

	r.Run()
}
