package router

import (
	"app/src/controller"
	"app/src/service"

	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(v1 fiber.Router, u service.AdminService) {
	// Initialize the UserController with the UserService
	adminController := controller.NewAdminController(u)
	// Define user-related routes
	adminGroup := v1.Group("/admins")
	adminGroup.Post("/", adminController.CreateAdmin)
	adminGroup.Get("/", adminController.GetAll)

}
