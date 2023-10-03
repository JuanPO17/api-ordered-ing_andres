package handlers

import (
	"github.com/JuanPO17/fiber-crud-api/database"
	"github.com/JuanPO17/fiber-crud-api/models"
	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello sebas")
}

func GetTasks(c *fiber.Ctx) error {
	db := database.DB
	var tasks []models.Task
	db.Find(&tasks)
	return c.JSON(tasks)
}

func GetSoloTasks(c *fiber.Ctx) error {
	id :=c.Params("id")
	db := database.DB
	var task models.Task
	db.Find(&task, id)
	return c.JSON(task)
}

func CreateTask(c *fiber.Ctx) error {
	db := database.DB
	task := new(models.Task)

	if err := c.BodyParser(task); err != nil {
		return c.Status(400).JSON(err)
	}
	
	db.Create(&task)
	return c.JSON(task)
}

func EditTask(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	
	var task models.Task
	if err := db.First(&task, id).Error; err != nil {
		return c.Status(404).JSON(err)
	}
	
	updateTask := new(models.Task)
	if err := c.BodyParser(updateTask); err != nil {
		return c.Status(400).JSON(err)
	}

	task.Body = updateTask.Body

	db.Save(&task)

	return c.JSON(task)

}

func DeleteTasks(c *fiber.Ctx) error {
	id :=c.Params("id")
	db := database.DB
	var task models.Task

	if err := db.First(&task, id).Error;err != nil {
		return c.Status(404).JSON(err)
	}
	db.Delete(&task)
	return c.JSON(fiber.Map{"Status": "success", "message": "Task Deleted", "id": id})
}