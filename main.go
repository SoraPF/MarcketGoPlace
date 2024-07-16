package main

import (
	"Marcketplace/config"
	"Marcketplace/controller"
	"Marcketplace/helper"
	"Marcketplace/repository"
	"Marcketplace/router"
	"Marcketplace/services"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)

func main() {
	fmt.Print("Run service ...")

	loagConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to connect database", err)
	}

	db := config.ConnectionDB(&loagConfig)
	err = config.AutoMigrate(db)
	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}

	fmt.Println("Database migration successful")

	fmt.Println("loading needs for templates")

	// Initialisation du moteur de templates
	engine := django.New("./views", ".django")
	engine.Reload(true)
	engine.Debug(true)

	engine.AddFunc("lower", helper.ToLower)

	// Créer une nouvelle instance de l'application Fiber en utilisant le moteur de templates
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/public", "./public")

	// Routes pour les templates
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	fmt.Println("loading needs for routes and crud for user and items")
	validate := validator.New()

	ObjRepo := repository.NewObjRepositoryImpl(db)
	ObjServ := services.NewObjServiceImpl(ObjRepo, validate)
	objCon := controller.NewObjController(ObjServ)
	objRoutes := router.ObjRoute(objCon)

	userRpo := repository.NewuserRepositoryImpl(db)
	userServ := services.NewUserServiceImpl(userRpo, validate)
	userCon := controller.NewuserController(userServ)
	userRoutes := router.UserRoute(userCon)

	// Grouper les routes de l'API sous le préfixe "/api"
	app.Mount("/", router.AuthentRoutes(userCon))

	api := app.Group("/api")
	api.Mount("/", objRoutes)
	api.Mount("/", userRoutes)
	api.Mount("/", router.Authentification(userCon))

	// Lancer l'application sur le port 3000
	log.Fatal(app.Listen(":4200"))
}
