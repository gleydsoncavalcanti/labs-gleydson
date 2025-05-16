package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/prea/go-api/models"
    "github.com/prea/go-api/database"
)

func GetBooks(c *gin.Context) {
    var books []models.Book
    database.DB.Find(&books)
    c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    database.DB.Create(&book)
    c.JSON(http.StatusOK, book)
}
