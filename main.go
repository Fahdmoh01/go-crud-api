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
	r.Run()
}
