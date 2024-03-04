package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rookie623/gin-crud/initializers"
	"github.com/rookie623/gin-crud/models"
	"net/http"
)

func PostCreate(c *gin.Context) {
	//Get data of req body
	var body struct {
		Body  string
		Title string
	}

	err := c.Bind(&body)
	if err != nil {
		return
	}

	//Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	//Return it
	c.JSON(http.StatusCreated, gin.H{
		"post": post,
	})

}

func PostIndex(c *gin.Context) {
	//Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond with them
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func PostGet(c *gin.Context) {
	//Get id of url
	id := c.Param("id")
	//Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	//Respond with the Post
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	//Get the id off the url
	id := c.Param("id")

	//Get the data off req body
	var post models.Post
	initializers.DB.First(&post, id)

	//Find the post were updating
	var body struct {
		Body  string
		Title string
	}

	err := c.Bind(&body)
	if err != nil {
		return
	}

	//Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	//Respond with it
	c.JSON(http.StatusAccepted, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	//Get the id off the url
	id := c.Param("id")

	//Delete it
	initializers.DB.Delete(&models.Post{}, id)

	//Respond
	c.Status(http.StatusOK)
}
