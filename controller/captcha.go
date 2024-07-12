package controller

import (
	"bytes"
	"image/png"

	"github.com/afocus/captcha"
	"github.com/gofiber/fiber/v2"
)

var cap *captcha.Captcha

func init() {
	cap = captcha.New()
	if err := cap.SetFont("Outwrite.ttf"); err != nil {
		panic(err)
	}
}

func Captcha(c *fiber.Ctx) error {
	img, str := cap.Create(6, captcha.NUM)
	c.Cookie(&fiber.Cookie{
		Name:  "captcha",
		Value: str,
		Path:  "/",
	})
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return err
	}
	c.Set("Content-Type", "image/png")
	return c.Send(buf.Bytes())
}
