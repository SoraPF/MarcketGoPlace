package router

import (
	"Marcketplace/controller"
	"Marcketplace/data/response"

	"github.com/gofiber/fiber/v2"
)

func AuthentRoutes(userController *controller.UserController, categories []response.CategoryResponse) *fiber.App {
	router := fiber.New()

	router.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("layouts/login", fiber.Map{
			"Title":      "LOGIN",
			"Categories": categories,
		})
	})

	router.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("layouts/register", fiber.Map{
			"Title":      "REGISTER",
			"Categories": categories,
		})
	})

	router.Get("/TFA-generate/:id", func(c *fiber.Ctx) error {
		return c.Render("layouts/generate2fa", fiber.Map{
			"Title": "double authentification generate",
		})
	})

	router.Get("/TFA-validate/:id", func(c *fiber.Ctx) error {
		return c.Render("layouts/validate2fa", fiber.Map{
			"Title": "double authentification validate",
		})
	})
	return router
}
