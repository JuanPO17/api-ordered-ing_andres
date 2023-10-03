package router

import (
	"github.com/JuanPO17/fiber-crud-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/hello", handlers.Hello)
	app.Get("/get", handlers.GetTasks)
	app.Post("/post", handlers.CreateTask)
	app.Get("/get/:id", handlers.GetSoloTasks)
	app.Put("/put/:id", handlers.EditTask)
	app.Delete("/del/:id", handlers.CreateTask)
}