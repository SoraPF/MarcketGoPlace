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
	"github.com/gofiber/contrib/websocket"
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
	engine.AddFunc("capitalize", helper.CapitalizeFirstLetter)

	// Créer une nouvelle instance de l'application Fiber en utilisant le moteur de templates
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/public", "./public")

	fmt.Println("loading needs for routes and crud for user and items")
	validate := validator.New()

	eRepo := repository.NewElemRepositoryImpl(db)
	eServ := services.NewEleServiceImpl(eRepo, validate)
	eCon := controller.NewElemController(eServ)

	ObjRepo := repository.NewObjRepositoryImpl(db)
	ObjServ := services.NewObjServiceImpl(ObjRepo, validate)
	objCon := controller.NewObjController(ObjServ)
	objRoutes := router.ObjRoute(objCon)

	userRpo := repository.NewuserRepositoryImpl(db)
	userServ := services.NewUserServiceImpl(userRpo, validate)
	userCon := controller.NewuserController(userServ)
	userRoutes := router.UserRoute(userCon)

	messRepo := repository.NewMessageRepositoryImpl(db)
	messServ := services.NewMesServiceImpl(messRepo, validate)
	messCon := controller.NewMessController(messServ)
	//objRoutes := router.ObjRoute(objCon)

	categories, err := eCon.GetCategories()
	if err != nil {
		log.Fatal("Failed to catch", err)
	}
	tags, err := eCon.GetTags()
	if err != nil {
		log.Fatal("Failed to catch", err)
	}

	// Routes pour les templates
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title":      "Marchet Go Place!",
			"Categories": categories,
		})
	})
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.Render("main", fiber.Map{
			"Title":      "Test!",
			"Categories": categories,
		})
	})

	// Grouper les routes
	app.Get("/ws", websocket.New(controller.HandleConnections))
	go controller.HandleMessages()
	app.Mount("/", router.AuthentRoutes(userCon, categories))
	app.Mount("/", router.Aouth2())
	app.Mount("/", router.Robject(objCon, categories, tags))
	app.Mount("/", router.FrontMessenger(messCon, categories))

	api := app.Group("/api")
	api.Mount("/", objRoutes)
	api.Mount("/", userRoutes)
	api.Mount("/", router.BackendRoutes(userCon, objCon, messCon))

	// Lancer l'application sur le port 3000
	log.Fatal(app.Listen(":3000"))
}
