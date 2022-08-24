package main

import (
	"github.com/gin-gonic/gin"
	db "user-management/adapter/databases"
	"user-management/api/routes"
)

func main() {
	db.Connect()
	db.AutoMigrate()
	server := gin.Default()

	routes.Setup(server)

	server.Run(":8000")
}
