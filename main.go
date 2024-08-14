package main

import (
	"github.com/thiagocprado/golang-api-rest-gin/database"
	"github.com/thiagocprado/golang-api-rest-gin/routes"
)

func main() {
	database.ConnectDB()
	routes.HandleRequests()
}
