package handlers

// -----------------------------gestion token et gestion authentification---------------------
/*

	claims := jwt.MapClaims{
		"email": email,
		"admin": false,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": t, "uid": user.Id, "profession": user.Profession})


func IsLogin(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	if tokenString == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("méthode de signature non valide")
		}
		return []byte("secret"), nil
	})

	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetGenerate2fa(c *fiber.Ctx) error {
	u := data.NewUser()
	if err := c.BodyParser(u); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Le body est vide")
	}

	user := u.FindUser()

	if user.IsNFA() {
		println("not generate", user.NFA.Secret)
		return c.JSON(fiber.Map{
			"secret": user.NFA.Secret,
			"url":    user.NFA.URL,
			"qr":     user.NFA.QR,
		})
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "ticktingApp",
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
	user.NFA.URL = key.URL()
	user.NFA.QR = fmt.Sprintf("data:image/png;base64,%s", qrCodeBase64)
	user.UpdateUsers()
	println("generate", user.NFA.Secret)
	return c.JSON(fiber.Map{
		"secret": user.NFA.Secret,
		"url":    user.NFA.URL,
		"qr":     user.NFA.QR,
	})
}

func GetValidate2fa(c *fiber.Ctx) error {
	type Request struct {
		Code   string `json:"code"`
		Secret string `json:"secret"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	valid := totp.Validate(req.Code, req.Secret)
	if valid {
		return c.SendString("2FA code is valid")
	} else {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid 2FA code")
	}
}
*/
