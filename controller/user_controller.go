package controller

import (
	"Marcketplace/data/request"
	"Marcketplace/data/response"
	"Marcketplace/helper"
	"Marcketplace/services"
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/pquerna/otp/totp"
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

	captchaValue := ctx.Cookies("captcha")
	if captchaValue == "" || createUserRequest.Captcha != captchaValue {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid CAPTCHA",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
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

	captchaValue := ctx.Cookies("captcha")
	if captchaValue == "" || req.Captcha != captchaValue {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid CAPTCHA",
		})
	}

	valid, user, err := uc.UserService.AuthenticateUser(req.Email, req.Password)
	if err != nil || !valid {
		return ctx.Status(fiber.StatusUnauthorized).Render("layouts/login", fiber.Map{
			"Title": "Login",
			"Error": "Invalid email or password",
		})
	}

	if user.IsNFA() {
		// Rediriger vers la page de validation 2FA
		webResponse := map[string]interface{}{
			"code":         200,
			"status":       "ok",
			"message":      "2FA required",
			"redirect_url": "/validate-2fa", // URL de redirection pour la validation 2FA
		}
		return ctx.Status(fiber.StatusOK).JSON(webResponse)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
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
		"token":        tokenString, // Notez qu'on retourne le token sous forme de string
		"redirect_url": "/",         // URL de redirection après connexion
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (uc *UserController) GetGenerate2FA(c *fiber.Ctx) error {
	UserId := c.Params("id")
	id, err := strconv.Atoi(UserId)
	helper.ErrorPanic(err)

	user := uc.UserService.FindUser(id)

	if services.IsNFA(user) {
		return c.JSON(fiber.Map{
			"secret": user.NFA.Secret,
			"qr":     user.NFA.QRcode,
		})
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "MarketplaceApp",
		AccountName: user.Email,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error generating TOTP secret")
	}

	image, err := key.Image(200, 200)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error generating QR code")
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, image); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error encoding QR code image")
	}
	qrCodeBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())

	user.NFA.Secret = key.Secret()
	user.NFA.QRcode = fmt.Sprintf("data:image/png;base64,%s", qrCodeBase64)
	// Assuming you have a method to update user
	uc.UserUpdate(c)

	return c.JSON(fiber.Map{
		"secret": user.NFA.Secret,
		"qr":     user.NFA.QRcode,
	})
}

func (uc *UserController) GetValidate2FA(c *fiber.Ctx) error {
	type Request struct {
		Code   string `json:"code"`
		UserId int    `json:"user_id"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	user := uc.UserService.FindUser(req.UserId)
	if user == nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	valid := totp.Validate(req.Code, user.NFA.Secret)
	if valid {
		return c.SendString("2FA code is valid")
	} else {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid 2FA code")
	}
}
