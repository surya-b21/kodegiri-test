package router

import (
	"kodegiri/app/controller/auth"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Handle(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/info", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(
			fiber.Map{
				"status":  200,
				"message": "info api",
			},
		)
	})

	// auth route
	api.Post("/login", auth.Login)

	protected := api.Group("", jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("s3cret!!443")},
	}))

	protected.Get("/protected", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(
			fiber.Map{
				"status":  200,
				"message": "restricted api",
			},
		)
	})
}
