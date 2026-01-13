package controller

import (
	"app/src/service"
	"app/src/validation"

	"github.com/gofiber/fiber/v2"
)

type AdminController struct {
	_AdminService service.AdminService
}

func NewAdminController(adminService service.AdminService) *AdminController {
	return &AdminController{

		_AdminService: adminService,
	}
}

// Create
func (u *AdminController) CreateAdmin(c *fiber.Ctx) error {

	admin, err := u._AdminService.CreateAdmin(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(admin)
}

// Get All
func (u *AdminController) GetAll(c *fiber.Ctx) error {
	query := &validation.QueryAdmin{

		Page:   c.QueryInt("Page", 1),
		Limit:  c.QueryInt("Limit", 20),
		Search: c.Query("Search", ""),
	}

	admins, err := u._AdminService.GetAll(c, query)

	if err != nil {

		return err
	}
	return c.Status(fiber.StatusOK).JSON(admins)
}
