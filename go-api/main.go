package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gleydson/go-api/routes"
    "github.com/gleydson/go-api/database"
)

func main() {
    database.ConnectDB()
    r := gin.Default()
    routes.SetupRoutes(r)
    r.Run(":8080")
}
