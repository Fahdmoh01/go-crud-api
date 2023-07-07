package main

import (
	"crud-api/initializers"
	"crud-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

//This create our table int he database
func main() {
	// Migrate the schema
	initializers.DB.AutoMigrate(&models.Post{})
}
