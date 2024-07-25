package router

import (
	"Marcketplace/controller"
	"Marcketplace/data/response"
	"strconv"

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

func Robject(ObjController *controller.ObjController, categories []response.CategoryResponse) *fiber.App {
	router := fiber.New()

	router.Get("/categories/:id", func(c *fiber.Ctx) error {
		CID := c.Params("id")
		id, err := strconv.Atoi(CID)
		if err != nil {
			return c.Render("categories", fiber.Map{
				"Title":      "categorie",
				"Categories": categories,
			})
		}
		cid := uint(id)
		articles := ObjController.ObjByCategID(cid)
		if articles == nil {
			return c.Render("categories", fiber.Map{
				"Title":      "categorie",
				"Categories": categories,
			})
		}
		return c.Render("categories", fiber.Map{
			"Title":      "categorie",
			"Categories": categories,
			"articles":   articles,
		})
	})

	router.Get("/article/:id", func(c *fiber.Ctx) error {
		CID := c.Params("id")
		id, err := strconv.Atoi(CID)
		if err != nil {
			return c.Render("article", fiber.Map{
				"Title":      "categorie",
				"Categories": categories,
			})
		}
		cid := uint(id)
		article, err := ObjController.ObjByArticleID(cid)
		if err != nil {
			return c.Render("article", fiber.Map{
				"Title":      "categorie",
				"Categories": categories,
			})
		}
		return c.Render("article", fiber.Map{
			"Title":      "categorie",
			"Categories": categories,
			"article":    article,
		})
	})

	return router
}
