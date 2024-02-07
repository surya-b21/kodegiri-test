package router

import (
	"kodegiri/app/controller/auth"
	"kodegiri/app/controller/tiermanagement"

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

	// tier management route
	protected.Get("/tier-management", tiermanagement.ListTierManagement)
	protected.Post("/tier-management", tiermanagement.StoreTierManagement)
	protected.Get("/tier-management/:id", tiermanagement.DetailTierManagement)
	protected.Put("/tier-management/:id", tiermanagement.UpdateTierManagement)
	protected.Delete("/tier-management/:id", tiermanagement.DeleteTierManagement)
}
