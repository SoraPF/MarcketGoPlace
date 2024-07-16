package router

import (
	"Marcketplace/config"

	"github.com/gofiber/fiber/v2"
)

func Aouth2() *fiber.App {
	router := fiber.New()

	router.Get("/login", config.HandleGoogleLogin)
	router.Get("/callback", config.HandleGoogleCallback)

	return router
}
