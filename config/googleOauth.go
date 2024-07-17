package config

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	GoogleOauthConfig *oauth2.Config
	OauthStateString  = "random"
)

func init() {
	err := godotenv.Load("oauth.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	GoogleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:3000/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func HandleGoogleLogin(c *fiber.Ctx) error {
	url := GoogleOauthConfig.AuthCodeURL(OauthStateString)
	return c.Redirect(url)
}

func HandleGoogleCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != OauthStateString {
		return c.Status(http.StatusUnauthorized).SendString("State is invalid")
	}

	code := c.Query("code")
	token, err := GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Could not get token")
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Could not create request")
	}
	defer response.Body.Close()

	userInfo := struct {
		ID      string `json:"id"`
		Email   string `json:"email"`
		Picture string `json:"picture"`
	}{}
	if err := json.NewDecoder(response.Body).Decode(&userInfo); err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Could not parse response")
	}

	return c.JSON(userInfo)
}
