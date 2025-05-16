package database

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/prea/go-api/models"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "host=localhost user=postgres password=postgres dbname=go_api port=5432 sslmode=disable"
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Falha ao conectar ao banco de dados!")
    }
    database.AutoMigrate(&models.Book{})
    DB = database
}
