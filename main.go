package main

import (
	"copy_crud_api/config"
	"copy_crud_api/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	//call connect database function
	//returning db, err
	db, err := config.DatabaseConnection()
	if err != nil {
		log.Fatalf("database connection error: %v", err)
	}
	defer db.Close()

	//creating default gin router for logger & recovery
	router := gin.Default()
	routes.RegisterRoutes(router, db) //also send db with router

	port := os.Getenv("PORT")
	router.Run(":" + port)
}
