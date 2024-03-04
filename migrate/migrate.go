package main

import (
	"github.com/rookie623/gin-crud/initializers"
	"github.com/rookie623/gin-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Post{})
	if err != nil {
		return
	}
}
