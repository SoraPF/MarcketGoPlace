package controller

import (
	"Marcketplace/data/request"
	"Marcketplace/data/response"
	"Marcketplace/helper"
	"Marcketplace/model/entities"
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
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
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

	UserId := ctx.Params("id")
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
	UserId := ctx.Params("id")
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
	UserId := ctx.Params("id")
	id, err := strconv.Atoi(UserId)
	helper.ErrorPanic(err)
	uid := uint(id)
	UserController := controller.UserService.FindById(uid)

	if UserController.Email == "" {
		return ctx.Status(fiber.StatusNotFound).SendString("utilisateur pas trouver")
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully delete Users data!",
		Data:    UserController,
	}
	return ctx.Status(fiber.StatusFound).JSON(webResponse)
}

func (controller *UserController) UserFindAll(ctx *fiber.Ctx) error {
	UserController := controller.UserService.FindAll()

	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "Successfully delete Users data!",
		Data:    UserController,
	}
	return ctx.Status(fiber.StatusFound).JSON(webResponse)
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

	captchaValue := ctx.Cookies("captcha")
	if captchaValue == "" || req.Captcha != captchaValue {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid CAPTCHA",
		})
	}

	if user.IsNFA() {
		webResponse := map[string]interface{}{
			"code":         200,
			"status":       "ok",
			"message":      "2FA required",
			"redirect_url": fmt.Sprintf("/TFA-validate/%d", user.Id),
		}
		return ctx.Status(fiber.StatusOK).JSON(webResponse)
	}

	token := CreateToken(*user)

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).Render("layouts/login", fiber.Map{
			"Title": "Login",
			"Error": "Internal server error",
		})
	}

	userIDStr := strconv.FormatUint(uint64(user.Id), 10)

	ctx.Cookie(&fiber.Cookie{
		Name:    "jwt-" + string(userIDStr),
		Value:   tokenString,
		Expires: time.Now().Add(time.Hour * 1),
	})

	ctx.Cookie(&fiber.Cookie{
		Name:    "user_id",
		Value:   fmt.Sprintf("%d", user.Id),
		Expires: time.Now().Add(time.Hour * 1),
	})

	webResponse := map[string]interface{}{
		"code":         200,
		"status":       "ok",
		"message":      "Login successful!",
		"token":        tokenString, // Notez qu'on retourne le token sous forme de string
		"redirect_url": "/",         // URL de redirection après connexion
	}
	println("login")
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (uc *UserController) GetGenerate2FA(c *fiber.Ctx) error {
	userId := c.Params("id")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	user := uc.UserService.FindUser(id)
	if user == nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	if services.IsNFA(user) {
		nfa, err := uc.UserService.FindNFA(user.NFAID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("internal corruption whene searching QRcode for double authentification")
		}
		return c.JSON(fiber.Map{
			"secret": nfa.Secret,
			"qr":     nfa.QRcode,
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

	nfa := &entities.NFA{
		Secret: key.Secret(),
		QRcode: fmt.Sprintf("data:image/png;base64,%s", qrCodeBase64),
	}
	err = uc.UserService.CreateNFA(nfa)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating NFA")
	}

	user.NFAID = &nfa.ID
	updateRequest := request.UpdateUserRequest{
		ID:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password, // Keep the existing password
		NFAID:    &nfa.ID,
	}
	uc.UserService.Update(updateRequest)

	return c.JSON(fiber.Map{
		"secret": nfa.Secret,
		"qr":     nfa.QRcode,
	})
}
func (uc *UserController) GetValidate2FA(c *fiber.Ctx) error {
	type Request struct {
		Code   string `json:"code"`
		UserId int    `json:"userId"`
	}
	userId := c.Params("id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		fmt.Println("Invalid user ID:", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		fmt.Println("Body parsing error:", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}
	req.UserId = id

	user := uc.UserService.FindUser(req.UserId)
	if user == nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	if user.NFAID == nil {
		return c.Status(fiber.StatusNotFound).SendString("NFA not found for user")
	}

	nfa, err := uc.UserService.FindNFA(user.NFAID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("NFA not found")
	}

	valid := totp.Validate(req.Code, nfa.Secret)
	if !valid {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid 2FA code")
	} else {
		token := CreateToken(*user)

		tokenString, err := token.SignedString([]byte("secret"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).Render("layouts/login", fiber.Map{
				"Title": "Login",
				"Error": "Internal server error",
			})
		}

		c.Cookie(&fiber.Cookie{
			Name:    "jwt-" + string(user.ID),
			Value:   tokenString,
			Expires: time.Now().Add(time.Hour * 1),
		})

		webResponse := map[string]interface{}{
			"code":         200,
			"status":       "ok",
			"message":      "Login successful!",
			"token":        tokenString,
			"redirect_url": "/",
		}
		return c.Status(fiber.StatusOK).JSON(webResponse)
	}
}

func IsLogin(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	println("token:", token)
	_, err := VerifyToken(token)
	if err != nil {
		println("error:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	userIDStr := c.Cookies("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	return c.Status(fiber.StatusOK).JSON(userID)
}

func (uc *UserController) Logout(c *fiber.Ctx) error {
	type Logout struct {
		UserId string `json:"userID"`
	}
	var user Logout
	if err := c.BodyParser(&user); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	println(user.UserId)
	uid := user.UserId
	c.Cookie(&fiber.Cookie{
		Name:    "captcha",
		Expires: time.Unix(0, 0),
		MaxAge:  -1,
	})
	c.Cookie(&fiber.Cookie{
		Name:    "jwt-" + uid,
		Expires: time.Unix(0, 0),
		MaxAge:  -1,
	})
	c.Cookie(&fiber.Cookie{
		Name:    "user_id",
		Expires: time.Unix(0, 0),
		MaxAge:  -1,
	})
	return c.SendStatus(fiber.StatusOK)
}
