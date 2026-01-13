package router

import "github.com/gofiber/fiber/v2"

func FrontendRoutes(app *fiber.App) {
	app.Static("/frontend", "./frontend")
	app.Static("/uploads", "./frontend/uploads")

	// Protected route helper
	protected := func(c *fiber.Ctx) error {
		if c.Cookies("auth_token") == "" {
			return c.Redirect("/login")
		}
		return c.Next()
	}

	app.Get("/", func(c *fiber.Ctx) error { return c.SendFile("./frontend/index.html") })

	app.Get("/login", func(c *fiber.Ctx) error {
		if c.Cookies("auth_token") != "" {
			return c.Redirect("/dashboard")
		}
		return c.SendFile("./frontend/login.html")
	})

	app.Get("/admin", protected, func(c *fiber.Ctx) error { return c.SendFile("./frontend/admin.html") })
	app.Get("/dashboard", protected, func(c *fiber.Ctx) error { return c.SendFile("./frontend/dashboard.html") })
}
