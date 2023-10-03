package main

import (
	"log"

	"github.com/JuanPO17/fiber-crud-api/database"
	"github.com/JuanPO17/fiber-crud-api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)


func main() {
	database.ConnectDB()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://localhost:5173",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowCredentials: true,
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.SetupRoutes(app)
	
	log.Fatal(app.Listen(":8000"))
}

