package router

import (
	"Marcketplace/controller"

	"github.com/gofiber/fiber/v2"
)

func BackendRoutes(userController *controller.UserController, ObjController *controller.ObjController, messController *controller.MessageController) *fiber.App {
	router := fiber.New()

	router.Route("/authent", func(router fiber.Router) {
		router.Post("/login", userController.Login)
		router.Post("/register", userController.UserCreate)
		router.Get("/isLogin", controller.IsLogin)
		router.Get("/logout", userController.Logout)
	})

	router.Get("/captcha", controller.Captcha)
	router.Get("/generate-2fa/:id", userController.GetGenerate2FA)
	router.Post("/validate-2fa/:id", userController.GetValidate2FA)
	router.Post("/articles/verify", ObjController.AdminResponceNewArticle) //presque bon
	router.Post("/article/by-name", ObjController.FindByName)

	router.Route("/messenger", func(router fiber.Router) {
		router.Post("/create", messController.CreateConversation)
		router.Post("/delete", messController.SupprimerConversation)
		router.Post("/newMesage", messController.SendMessage)
		router.Post("/check", messController.CheckMessenger)

		//router.Get("/getmessages/:id", messController.GetMessagesFromConversation)
	})

	return router
}
