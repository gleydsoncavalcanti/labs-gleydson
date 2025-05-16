package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/prea/go-api/controllers"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/books", controllers.GetBooks)
    r.POST("/books", controllers.CreateBook)
}
