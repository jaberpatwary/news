package router

import (
	"app/src/controller"
	"app/src/service"

	"github.com/gofiber/fiber/v2"
)

func ComRoutes(v1 fiber.Router, u service.ComService) {
	// Initialize the UserController with the UserService
	comController := controller.NewComController(u)
	// Define user-related routes
	comGroup := v1.Group("/coms")
	comGroup.Post("/", comController.CreateCom)
	comGroup.Get("/", comController.GetAll)

}
