package database

import (
	"log"

	"github.com/thiagocprado/golang-api-rest-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados!", err)
	}

	DB.AutoMigrate(&models.Student{})
}
