package auth

import (
	"crypto/rand"
	"encoding/base64"
	"kodegiri/app/model"
	"kodegiri/app/services"
	"net/mail"
	"time"

	"github.com/dlclark/regexp2"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginPayload struct {
	Email         *string `json:"email,omitempty"`
	Password      *string `json:"password,omitempty"`
	RememberToken *string `json:"remember_token,omitempty"`
}

// Login godoc
// @Summary      Login function
// @Description  Get token for auth
// @Tags         Auth
// @Accept       json
// @Produce		 json
// @Param        body body LoginPayload true "Login Payload"
// @Success      200
// @Router       /login [post]
func Login(c *fiber.Ctx) error {
	db := services.DB

	payload := LoginPayload{}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Payload",
		})
	}

	//validate email
	if payload.Email != nil {
		if _, err := mail.ParseAddress(*payload.Email); nil != err {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid email",
			})
		}
	}

	// validate password
	if payload.Password != nil {
		pattern := `^(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9])(?=.*[!@#$%^&*()_+=-])[A-Za-z0-9!@#$%^&*()_+=-]{8,}$`
		regex := regexp2.MustCompile(pattern, 0)
		if match, _ := regex.MatchString(*payload.Password); !match {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid password",
			})
		}
	}

	// find user
	user := model.User{}
	db.Where("email = ?", payload.Email).Or("remember_token = ?", payload.RememberToken).First(&user)

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
		db.Model(&user).Update("remember_token", user.RememberToken)
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
		"expiry":         time.Now().Add(time.Hour * 72).UnixMilli(),
	})
}
