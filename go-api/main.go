package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "go-api/database"
    "go-api/routes"
)

func main() {
    database.ConnectDB()

    router := gin.Default()

    router.StaticFile("/", "./static/index.html")

    router.Static("/static", "./static")

    // Define rotas da API
    routes.SetupRoutes(router)

    port := getPort()
    log.Printf("Servidor rodando em http://localhost:%s", port)
    if err := router.Run(":" + port); err != nil {
        log.Fatalf("Erro ao iniciar servidor: %v", err)
    }
}

func getPort() string {
    if port := os.Getenv("PORT"); port != "" {
        return port
    }
    return "8080"
}
