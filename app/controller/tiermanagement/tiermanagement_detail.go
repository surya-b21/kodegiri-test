package tiermanagement

import (
	"kodegiri/app/model"
	"kodegiri/app/services"

	"github.com/gofiber/fiber/v2"
)

// DetailTierManagemet
// @Summary      Tier Detail function
// @Description  Get a tier detail
// @Tags TierManagement
// @Security ApiKeyAuth
// @Param id path string true "Tier id"
// @Success 200 {object} model.Tier
// @Failure 404
// @Router       /tier-management/{id} [get]
func DetailTierManagement(c *fiber.Ctx) error {
	db := services.DB

	tier := model.Tier{}
	mod := db.First(&tier, "id = ?", c.Params("id"))

	if mod.RowsAffected < 1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "tier not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(tier)
}
