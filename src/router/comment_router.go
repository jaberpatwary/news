package router

import (
	"app/src/controller"
	"app/src/service"

	"github.com/gofiber/fiber/v2"
)

func CommentRoutes(v1 fiber.Router, u service.CommentService) {
	// Initialize the UserController with the UserService
	commentController := controller.NewCommentController(u)
	// Define user-related routes
	commentGroup := v1.Group("/comments")
	commentGroup.Post("/", commentController.CreateComment)
	commentGroup.Get("/", commentController.GetAll)
	commentGroup.Get("/:commentId", commentController.GetByCommentId)
	commentGroup.Put("/:commentId", commentController.UpdateComment)
	commentGroup.Delete("/:commentId", commentController.DeleteComment)
}
