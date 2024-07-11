package router

import (
	"Marcketplace/controller"

	"github.com/gofiber/fiber/v2"
)

func AuthentRoutes(userController *controller.UserController) *fiber.App {
	router := fiber.New()

	router.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("layouts/login", fiber.Map{
			"Title": "LOGIN",
		})
	})

	router.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("layouts/register", fiber.Map{
			"Title": "REGISTER",
		})
	})

	return router
}
