package tiermanagement

import (
	"kodegiri/app/model"
	"kodegiri/app/services"

	"github.com/gofiber/fiber/v2"
)

// StoreTierManagement func
// @Summary      Tier Store function
// @Description  Store a tier
// @Tags TierManagement
// @Params body json model.TierAPI true "Tier payload"
// @Success 200 {object} model.Tier
// @Failure 400
func StoreTierManagement(c *fiber.Ctx) error {
	payload := model.TierAPI{}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Payload",
		})
	}

	db := services.DB

	tier := model.Tier{}
	tier.Name = payload.Name
	tier.MaximalPoint = payload.MaximalPoint
	tier.MinimalPoint = payload.MinimalPoint
	db.Create(&tier)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Successfully creaated new tier",
		"tier":    tier,
	})
}
