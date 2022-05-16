package controllers

import (
	"net/http"

	"github.com/curtisvermeeren/go-rest-api-with-gin-and-gorm/models"
	"github.com/gin-gonic/gin"
)

// GET /books
// Returns all books as a JSON response
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// POST /books
// Create a new book
func CreateBook(c *gin.Context) {
	// Validate input from post
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	// Return the created book as JSON
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// GET /books/:id
// Returns a book based on id as a JSON response
func FindBook(c *gin.Context) {
	var book models.Book

	// Check the DB for an entry with matching id and return the first matching instance
	err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error
	if err != nil {
		// Return error a json error if the book was not found
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	// Return the book as JSON
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// PATCH /books/:id
// Updates a book with the matching id
func UpdateBook(c *gin.Context) {

	// Find if the book exists in the db
	var book models.Book
	err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error
	if err != nil {
		// Return error a json error if the book was not found
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	// Validate the new input
	var input models.UpdateBookInput
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Apply updates to the book
	models.DB.Model(&book).Updates(input)

	// Return the new book
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /books/:id
// Delete the book with the matching id
func DeleteBook(c *gin.Context) {
	// Find if the book exists in the db
	var book models.Book
	err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error
	if err != nil {
		// Return error a json error if the book was not found
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
