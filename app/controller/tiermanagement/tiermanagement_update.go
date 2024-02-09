package tiermanagement

import (
	"kodegiri/app/model"
	"kodegiri/app/services"

	"github.com/gofiber/fiber/v2"
)

// UpdateTierManagement func
// @Summary      Tier Update function
// @Description  Update a tier
// @Security ApiKeyAuth
// @Tags TierManagement
// @Param        body body model.TierAPI true "Member Activity Payload"
// @Success 200 {object} model.Tier
// @Failure 400
// @Router       /tier-management/{id} [put]
func UpdateTierManagement(c *fiber.Ctx) error {
	payload := model.TierAPI{}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Payload",
		})
	}

	db := services.DB

	tier := model.Tier{}
	mod := db.First(&tier, "id = ?", c.Params("id"))

	if mod.RowsAffected < 1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "tier not found",
		})
	}

	tier.Name = payload.Name
	tier.MaximalPoint = payload.MaximalPoint
	tier.MinimalPoint = payload.MinimalPoint
	db.Save(&tier)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully update tier",
		"tier":    tier,
	})
}
