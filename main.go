package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/santhosh1188/employee-api/database"
	"github.com/santhosh1188/employee-api/models"
	"github.com/santhosh1188/employee-api/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	database.Connect()

	database.DB.AutoMigrate(&models.Employee{})

	router := routes.RegisterEmployeeRoutes()

	fmt.Println("🚀 Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
