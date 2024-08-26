package router

import (
	"Marcketplace/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRouteur(noteController *controller.NoteController) *fiber.App {
	router := fiber.New()

	router.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "welcom to Golang, Fiber, and gorm",
		})
	})

	router.Route("/notes", func(router fiber.Router) {
		router.Post("/", noteController.Create)
		router.Get("/list", noteController.FindAll)
		router.Route("/:noteId", func(router fiber.Router) {
			router.Delete("/delete", noteController.Delete)
			router.Get("/find", noteController.FindById)
			router.Patch("/update", noteController.Update)
		})
	})

	return router
}

func ObjRoute(ObjController *controller.ObjController) *fiber.App {
	router := fiber.New()
	router.Get("/healthcheckerObj", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "welcom to Golang, Fiber, and gorm",
		})
	})
	router.Route("/objets", func(router fiber.Router) {
		router.Post("/create", ObjController.ObjCreate)
		router.Get("/list", ObjController.ObjFindAll)
		router.Route("/:objId", func(router fiber.Router) {
			router.Delete("/delete", ObjController.ObjDelete)
			router.Get("/find", ObjController.ObjFindById)
			router.Patch("/update", ObjController.ObjUpdate)
		})
	})
	return router
}

func UserRoute(userController *controller.UserController) *fiber.App {
	router := fiber.New()
	router.Get("/healthcheckerUser", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "welcom to Golang, Fiber, and gorm",
		})
	})
	router.Route("/user", func(router fiber.Router) {
		router.Post("/create", userController.UserCreate)
		router.Get("/list", userController.UserFindAll)
		router.Route("/:id", func(router fiber.Router) {
			router.Delete("/delete", userController.UserDelete)
			router.Get("/find", userController.UserFindById)
			router.Patch("/update", userController.UserUpdate)
		})
	})
	return router
}
