package controllers

import (
	"crud-api/initializers"
	"crud-api/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	post := models.Post{Title: "Fahd", Body: "Waguan baby, how's the body"}
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
