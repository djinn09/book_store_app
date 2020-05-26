package controllers

import (
	"crud/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	GET ALL books
*/
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

/*
{
  "title": "Start with",
  "author": "Simon Sink"
}
// POST /books
*/
func CreateBook(c *gin.Context) {
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindUsers(c *gin.Context) {
	var users []models.Users
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

/*
{
  "username": "Start with",
  "password": "Simon Sink"
}
// POST /users
*/
func CreateUser(c *gin.Context) {
	var input UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.Users{Username: input.Username, Password: input.Password}
	models.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}


// GET /books/:id
// Find a book
func FindBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

/*
{
  "title": "Start with",
  "author": "Simon Sink"
}
// PATCH /books
*/
func UpdateBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	fmt.Println("Herre")
	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&book).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

/*
// DELETE /books/:id
// Delete a book
*/
func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": "Success"})
}
