package database

import (
    "fmt"
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "go-api/models"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := getDSN()
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Erro ao conectar no banco de dados: %v", err)
    }

    err = db.AutoMigrate(&models.Book{})
    if err != nil {
        log.Fatalf("Erro ao fazer AutoMigrate: %v", err)
    }

    DB = db
    log.Println("Banco de dados conectado com sucesso.")
}

func getDSN() string {
    host := getEnv("DB_HOST", "localhost")
    user := getEnv("DB_USER", "postgres")
    password := getEnv("DB_PASSWORD", "postgres")
    dbname := getEnv("DB_NAME", "go_api")
    port := getEnv("DB_PORT", "5432")
    sslmode := getEnv("DB_SSLMODE", "disable")

    return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        host, user, password, dbname, port, sslmode)
}

func getEnv(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}
