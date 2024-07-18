package test

import (
	"Marcketplace/controller"
	"Marcketplace/data/request"
	"Marcketplace/data/response"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// UserServiceMock is a mock of the UserService interface

func TestCreateUserController(t *testing.T) {
	app := fiber.New()

	userServiceMock := new(UserServiceMock)
	userController := controller.NewuserController(userServiceMock)

	app.Get("/api/captcha", controller.Captcha)
	app.Post("/api/users/create", userController.UserCreate)

	// Step 1: Get the CAPTCHA
	captchaReq := httptest.NewRequest(http.MethodGet, "/api/captcha", nil)
	captchaResp, err := app.Test(captchaReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, captchaResp.StatusCode)

	// Read the CAPTCHA value from the response cookies
	var captchaValue string
	for _, cookie := range captchaResp.Cookies() {
		if cookie.Name == "captcha" {
			captchaValue = cookie.Value
			break
		}
	}
	assert.NotEmpty(t, captchaValue, "CAPTCHA value should not be empty")

	// Step 2: Create the user using the obtained CAPTCHA
	userRequest := request.CreateUserRequest{
		Username: "jason",
		Email:    "jason@email.com",
		Password: "jason",
		Captcha:  captchaValue,
	}

	userServiceMock.On("Create", userRequest).Return(nil)

	body, _ := json.Marshal(userRequest)
	req := httptest.NewRequest(http.MethodPost, "/api/users/create", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "captcha", Value: captchaValue})

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var webResponse response.Response
	err = json.NewDecoder(resp.Body).Decode(&webResponse)
	assert.NoError(t, err)
	assert.Equal(t, "ok", webResponse.Status)
	assert.Equal(t, "Successfully created user data!", webResponse.Message)

	userServiceMock.AssertExpectations(t)
}
