package controller

import (
	"app/src/service"
	"app/src/validation"

	"github.com/gofiber/fiber/v2"
)

type ComController struct {
	_ComService service.ComService
}

func NewComController(comService service.ComService) *ComController {
	return &ComController{

		_ComService: comService,
	}
}

// Create
func (u *ComController) CreateCom(c *fiber.Ctx) error {

	com, err := u._ComService.CreateCom(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(com)
}

// Get All
func (u *ComController) GetAll(c *fiber.Ctx) error {
	query := &validation.QueryCom{

		Page:   c.QueryInt("Page", 1),
		Limit:  c.QueryInt("Limit", 20),
		Search: c.Query("Search", ""),
	}

	coms, err := u._ComService.GetAll(c, query)

	if err != nil {

		return err
	}
	return c.Status(fiber.StatusOK).JSON(coms)
}
