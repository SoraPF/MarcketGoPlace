package controller

import (
	"Marcketplace/data/request"
	"Marcketplace/data/response"
	"Marcketplace/helper"
	"Marcketplace/services"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type UserController struct {
	UserService services.UserService
}

func NewuserController(service services.UserService) *UserController {
	return &UserController{UserService: service}
}

func (controller *UserController) UserCreate(ctx *fiber.Ctx) error {
	createUserRequest := request.CreateUserRequest{}
	if err := ctx.BodyParser(&createUserRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).Render("layouts/login", fiber.Map{
			"Title": "create",
			"Error": "Invalid request",
		})
	}

	controller.UserService.Create(createUserRequest)
	println("userCreated")

	valid, user, err := controller.UserService.AuthenticateUser(createUserRequest.Email, createUserRequest.Password)
	if err != nil || !valid {
		return ctx.Status(fiber.StatusUnauthorized).Render("layouts/login", fiber.Map{
			"Title": "create",
			"Error": "Internal probleme",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	webResponse := map[string]interface{}{
		"code":         200,
		"status":       "ok",
		"message":      "Login successful!",
		"token":        token,
		"redirect_url": "/", // URL de redirection après connexion
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *UserController) UserUpdate(ctx *fiber.Ctx) error {
	updateUserRequest := request.UpdateUserRequest{}
	err := ctx.BodyParser(&updateUserRequest)
	helper.ErrorPanic(err)

	UserId := ctx.Params("objId")
	id, err := strconv.ParseUint(UserId, 10, 32)
	helper.ErrorPanic(err)

	updateUserRequest.ID = uint(id)

	controller.UserService.Update(updateUserRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully update Users data!",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *UserController) UserDelete(ctx *fiber.Ctx) error {
	UserId := ctx.Params("UserId")
	id, err := strconv.Atoi(UserId)
	helper.ErrorPanic(err)
	controller.UserService.Delete(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully delete Users data!",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *UserController) UserFindById(ctx *fiber.Ctx) error {
	UserId := ctx.Params("objId")
	id, err := strconv.Atoi(UserId)
	helper.ErrorPanic(err)

	UserController := controller.UserService.FindById(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully delete Users data!",
		Data:    UserController,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *UserController) UserFindAll(ctx *fiber.Ctx) error {
	UserController := controller.UserService.FindAll()

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully delete Users data!",
		Data:    UserController,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (uc *UserController) Login(ctx *fiber.Ctx) error {
	req := request.LoginUser{}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).Render("layouts/login", fiber.Map{
			"Title": "Login",
			"Error": "Invalid request",
		})
	}

	valid, user, err := uc.UserService.AuthenticateUser(req.Email, req.Password)
	if err != nil || !valid {
		return ctx.Status(fiber.StatusUnauthorized).Render("layouts/login", fiber.Map{
			"Title": "Login",
			"Error": "Invalid email or password",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).Render("layouts/login", fiber.Map{
			"Title": "Login",
			"Error": "Internal server error",
		})
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "jwt",
		Value:   tokenString,
		Expires: time.Now().Add(time.Hour * 72),
	})
	println("giveNeeded")

	webResponse := map[string]interface{}{
		"code":         200,
		"status":       "ok",
		"message":      "Login successful!",
		"token":        token,
		"redirect_url": "/", // URL de redirection après connexion
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}
