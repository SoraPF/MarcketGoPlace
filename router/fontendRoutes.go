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

func Robject(ObjController *controller.ObjController, categories []response.CategoryResponse, tags []response.TagResponse) *fiber.App {
	router := fiber.New()

	router.Get("/categories/:id", func(c *fiber.Ctx) error {
		CID := c.Params("id")
		id, err := strconv.Atoi(CID)
		if err != nil {
			return c.Render("categories", fiber.Map{
				"Title":      "categorie",
				"Categories": categories,
				"tags":       tags,
			})
		}
		cid := uint(id)
		articles := ObjController.ObjByCategID(cid)
		if articles == nil {
			return c.Render("categories", fiber.Map{
				"Title":      "categorie",
				"Categories": categories,
				"tags":       tags,
			})
		}
		return c.Render("categories", fiber.Map{
			"Title":      "categorie",
			"Categories": categories,
			"articles":   articles,
			"tags":       tags,
		})
	})

	router.Get("/article/:id", func(c *fiber.Ctx) error {
		CID := c.Params("id")
		id, err := strconv.Atoi(CID)
		if err != nil {
			return c.Render("article", fiber.Map{
				"Title":      "article",
				"Categories": categories,
			})
		}
		cid := uint(id)
		article, err := ObjController.ObjByArticleID(cid)
		if err != nil {
			return c.Render("article", fiber.Map{
				"Title":      "article",
				"Categories": categories,
			})
		}
		println(article.Title)
		userID := c.Cookies("user_id")
		uid, err := strconv.Atoi(userID)
		if err != nil {
			println("bien mais pas bon")
			return c.Render("article", fiber.Map{
				"Title":      "article",
				"Categories": categories,
				"article":    article,
				"userID":     nil,
			})
		}
		println("bien mais pas bon")
		return c.Render("article", fiber.Map{
			"Title":      "article",
			"Categories": categories,
			"article":    article,
			"userID":     uid,
		})
	})

	router.Get("/article/search", func(c *fiber.Ctx) error {
		return c.Render("article", fiber.Map{
			"Title":      "categorie",
			"Categories": categories,
		})
	})

	router.Get("/createSell", func(c *fiber.Ctx) error {
		return c.Render("newArticle", fiber.Map{
			"Title":      "NEW ARTICLE",
			"Categories": categories,
			"tags":       tags,
		})
	})

	router.Get("/new-article/list", func(c *fiber.Ctx) error {
		articles, err := ObjController.GetArticles(0, "create")
		if err != nil || len(articles) == 0 {
			println("err1")
			return c.Render("VerifArticle", fiber.Map{
				"Title":      "verify article",
				"Categories": categories,
			})
		}
		return c.Render("VerifArticle", fiber.Map{
			"Title":      "List article",
			"Categories": categories,
			"Articles":   articles,
		})
	})

	router.Get("/new-article/:id", func(c *fiber.Ctx) error {
		CID := c.Params("id")
		id, err := strconv.Atoi(CID)
		if err != nil {
			println("err1")
			return c.Render("VerifArticle", fiber.Map{
				"Title":      "categorie",
				"Categories": categories,
			})
		}
		cid := uint(id)
		articles, err := ObjController.GetArticles(cid, "create")
		if err != nil {
			println("err2")
			return c.Render("VerifArticle", fiber.Map{
				"Title":      "verify article",
				"Categories": categories,
				"tags":       tags,
			})
		}
		return c.Render("VerifArticle", fiber.Map{
			"Title":      "verify article",
			"Categories": categories,
			"tags":       tags,
			"Article":    articles[0],
		})
	})

	router.Get("/createOk", func(c *fiber.Ctx) error {
		return c.Render("newArticle", fiber.Map{
			"Title":      "verify article",
			"Categories": categories,
		})
	})
	return router
}

func FrontMessenger(mc *controller.MessageController) *fiber.App {
	router := fiber.New()

	router.Get("/message/:id", func(c *fiber.Ctx) error {

		conv, err := mc.GetMessageFromConversation(c)
		if err != nil {
			erreur := "Il y a eu un problème lord de la vérification! veuillez réeseyer ultérieurement"
			return c.Render("messenger", fiber.Map{
				"Title":         "messenger",
				"messageErreur": erreur,
			})
		}
		StringuserID := c.Cookies("user_id")
		userID, err := strconv.Atoi(StringuserID)
		if err != nil {
			erreur := "Il y a eu un problème lord de la vérification! veuillez réeseyer ultérieurement"
			return c.Render("messenger", fiber.Map{
				"Title":         "messenger",
				"messageErreur": erreur,
			})
		}
		return c.Render("messenger", fiber.Map{
			"Title":    "messenger",
			"messages": conv,
			"userid":   userID,
		})
	})

	return router
}
