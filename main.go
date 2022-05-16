package main

import (
	"github.com/curtisvermeeren/go-rest-api-with-gin-and-gorm/controllers"
	"github.com/curtisvermeeren/go-rest-api-with-gin-and-gorm/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// Init a Gin router
	r := gin.Default()

	// Create a database connection
	models.ConnectDatabase()

	// Set routes
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	// Start the server using Run
	r.Run()
}
