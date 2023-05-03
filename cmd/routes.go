package main

import (
	"fmt"

	"github.com/Harshitha-git-hub/EmployeeAPI/database"
	"github.com/Harshitha-git-hub/EmployeeAPI/models"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/Employee", ListEmployee)
	app.Post("/Employee", CreateEmployee)
	app.Get("/Employee/:id", GetEmployeeById)
	app.Delete("/Employee/:id", DeleteEmployee)
	app.Put("/Employee/:id", UpdateEmployee)
}

func ListEmployee(c *fiber.Ctx) error {
	Employee := []models.Employee{}
	database.DB.Db.Find(&Employee)
	return c.Status(200).JSON(Employee)
}

func CreateEmployee(c *fiber.Ctx) error {
	Employee := new(models.Employee)
	if err := c.BodyParser(Employee); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&Employee)
	return c.Status(200).JSON(Employee)
}

func UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	var Employee models.Product

	database.DB.Db.Find(&Employee, "id = ?", id)

	if Employee.ID == ' ' {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	var updatedEmployee models.Employee

	err := c.BodyParser(&updatedEmployee)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	Employee.ID = updatedEmployee.ID
	Employee.Name = updatedEmployee.Name
	Employee.Salary = updatedEmployee.Salary
	Employee.Technologies = updatedEmployee.Technologies
	Employee.Projects = updatedEmployee.Projects
	Employee.Manager = updatedEmployee.Manager

	// Save the Changes
	database.DB.Db.Save(&Employee)
	return c.JSON(fiber.Map{"data": Employee})

}

func GetEmployeeById(c *fiber.Ctx) error {
	id := c.Params("id")
	var Employee models.Employee

	result := database.DB.Db.Find(&Employee, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&Employee)
}

func DeleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	var Employee models.Employee

	if result := database.DB.Db.First(&Employee, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	database.DB.Db.Delete(&Employee)

	return c.Status(200).JSON(&Employee)
}
