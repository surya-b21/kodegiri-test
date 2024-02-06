package auth

import (
	"crypto/rand"
	"encoding/base64"
	"kodegiri/app/model"
	"kodegiri/app/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginPayload struct {
	Email         *string `json:"email"`
	Password      *string `json:"password"`
	RememberToken *string `json:"remember_token"`
}

func Login(c *fiber.Ctx) error {
	db := services.DB

	payload := LoginPayload{}
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Payload",
		})
	}

	// find user
	user := model.Users{}
	db.Where("email = ?", *payload.Email).Or("remember_token = ?", *payload.RememberToken).First(&user)

	if user.Email == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if payload.RememberToken == nil {
		if err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(*payload.Password)); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Wrong password",
			})
		}
		b := make([]byte, 16)
		if _, err := rand.Read(b); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Something wrong happen",
			})
		}

		rememberToken := base64.URLEncoding.EncodeToString(b)
		user.RememberToken = &rememberToken
		db.Model(&user).Update("remember_token", *user.RememberToken)
	}

	claims := jwt.MapClaims{
		"name": *user.Name,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("s3cret!!443"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user":           user,
		"token":          t,
		"remember_token": user.RememberToken,
	})
}
