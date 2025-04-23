package main

import (
	"alle_assignment/app/routes"
	"alle_assignment/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDB()

	routes.SetupRoutes(r)

	r.Run(":8080") 
}
