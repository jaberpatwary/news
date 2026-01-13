package router

import (
	"app/src/config"
	"app/src/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Routes(app *fiber.App, db *gorm.DB) {

	UserService := service.NewUserService(db)
	CommentService := service.NewCommentService(db)
	ComService := service.NewComService(db)
	AdminService := service.NewAdminService(db)
	ArticleService := service.NewArticleService(db)

	v1 := app.Group("/v1")
	api := app.Group("/api")

	//HealthCheckRoutes(v1, healthCheckService)

	UserRoutes(v1, UserService)
	UserRoutes(api, UserService)
	CommentRoutes(v1, CommentService)
	CommentRoutes(api, CommentService)
	ComRoutes(v1, ComService)
	ComRoutes(api, ComService)
	ArticleRoutes(v1, ArticleService)  // Restful ones under v1
	ArticleRoutes(api, ArticleService) // New ones under /api for the new frontend
	AdminRoutes(v1, AdminService)
	AdminRoutes(api, AdminService)

	// TODO: add another routes here...

	if !config.IsProd {
		DocsRoutes(v1)
	}
}
