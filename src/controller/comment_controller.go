package controller

import (
	"app/src/service"
	"app/src/validation"

	"github.com/gofiber/fiber/v2"
)

type CommentController struct {
	_CommentService service.CommentService
}

func NewCommentController(commentService service.CommentService) *CommentController {
	return &CommentController{

		_CommentService: commentService,
	}
}

// Create
func (u *CommentController) CreateComment(c *fiber.Ctx) error {

	comment, err := u._CommentService.CreateComment(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(comment)
}

// Get All
func (u *CommentController) GetAll(c *fiber.Ctx) error {
	query := &validation.QueryComment{

		Page:   c.QueryInt("Page", 1),
		Limit:  c.QueryInt("Limit", 20),
		Search: c.Query("Search", ""),
	}

	comments, err := u._CommentService.GetAll(c, query)

	if err != nil {

		return err
	}
	return c.Status(fiber.StatusOK).JSON(comments)
}

// Get By Id
func (u *CommentController) GetByCommentId(c *fiber.Ctx) error {

	CommentId := c.Params("commentId")

	comment, err := u._CommentService.GetByCommentId(c, CommentId)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(comment)

}

// Update
func (u *CommentController) UpdateComment(c *fiber.Ctx) error {
	commentId := c.Params("commentId")

	var payload validation.UpdateComment

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON payload",
		})
	}

	// Service call
	updatedComment, err := u._CommentService.UpdateComment(c, commentId, &payload)
	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(updatedComment)
}

// delete
func (u *CommentController) DeleteComment(c *fiber.Ctx) error {
	commentID := c.Params("commentId")

	if commentID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"status":  "error",
			"message": "invalid comment id",
		})
	}

	if err := u._CommentService.DeleteComment(c, commentID); err != nil {

		if e, ok := err.(*fiber.Error); ok && e.Code == fiber.StatusNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"code":    fiber.StatusNotFound,
				"status":  "error",
				"message": "comment not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"status":  "success",
		"message": "Delete comment successfully",
	})
}
