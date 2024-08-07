package main

import (
	"github.com/thiagocprado/golang-api-rest-gin/database"
	"github.com/thiagocprado/golang-api-rest-gin/models"
	"github.com/thiagocprado/golang-api-rest-gin/routes"
)

func main() {
	database.ConnectDB()

	models.Students = []models.Student{
		{Name: "Thiago", CPF: "00000000000", RG: "123456789"},
		{Name: "Marcelo", CPF: "11111111111", RG: "987654321"},
	}

	routes.HandleRequests()
}
