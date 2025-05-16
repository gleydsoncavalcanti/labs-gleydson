package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "go-api/models"
    "go-api/database"
)

// GET /books
func GetBooks(c *gin.Context) {
    var books []models.Book
    database.DB.Find(&books)
    c.JSON(http.StatusOK, books)
}

// GET /books/:id
func GetBookByID(c *gin.Context) {
    id := c.Param("id")
    var book models.Book
    if err := database.DB.First(&book, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, book)
}

// POST /books
func CreateBook(c *gin.Context) {
    var input models.Book
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    book := models.Book{Title: input.Title, Author: input.Author}
    database.DB.Create(&book)
    c.JSON(http.StatusCreated, book)
}

// PUT /books/:id
func UpdateBook(c *gin.Context) {
    id := c.Param("id")

    var book models.Book
    if err := database.DB.First(&book, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    var input models.Book
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    book.Title = input.Title
    book.Author = input.Author

    database.DB.Save(&book)

    c.JSON(http.StatusOK, book)
}

// DELETE /books/:id
func DeleteBook(c *gin.Context) {
    id := c.Param("id")

    var book models.Book
    if err := database.DB.First(&book, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    database.DB.Delete(&book)

    c.JSON(http.StatusOK, gin.H{"message": "Book deleted", "id": id})
}

// OPTIONS /books
func OptionsBooks(c *gin.Context) {
    c.Header("Allow", "GET, POST, PUT, DELETE, OPTIONS")
    c.Status(http.StatusNoContent)
}
