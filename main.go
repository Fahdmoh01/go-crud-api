package main

import (
	"crud-api/controllers"
	"crud-api/initializers"

	"github.com/gin-gonic/gin"
)

//initialize env variables
func init() {
	initializers.LoadEnvVariables()
}

func main() {

	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)
	r.Run()
}
