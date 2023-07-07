package controllers

import (
	"crud-api/initializers"
	"crud-api/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostsCreate(c *gin.Context) {

	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	post := models.Post{
		Model: gorm.Model{},
		Title: body.Title,
		Body:  body.Body,
	}
	fmt.Printf("This is in POST %v", post)
	result := initializers.DB.Create(&post)
	fmt.Printf("This is in RESULT %v", result)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	//respond
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	//Get id from url
	id := c.Param("id")
	//Get the posts
	var post models.Post
	initializers.DB.First(&post, id)
	//respond
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostUpdate(c *gin.Context) {
	//Get id off url
	id := c.Param("id")

	//Get the data off the req body
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)
	//Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)
	//update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title, Body: body.Body,
	})

	//Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	//Get the id off the url
	id := c.Param("id")

	//Delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)
}
