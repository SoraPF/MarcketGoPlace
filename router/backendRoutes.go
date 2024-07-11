package router

import (
	"Marcketplace/controller"

	"github.com/gofiber/fiber/v2"
)

func Authentification(userController *controller.UserController) *fiber.App {
	router := fiber.New()

	router.Route("/authent", func(router fiber.Router) {
		router.Post("/login", userController.Login)
		router.Post("/register", userController.UserCreate)
		//router.Get("/isLogin")
	})

	return router
}
