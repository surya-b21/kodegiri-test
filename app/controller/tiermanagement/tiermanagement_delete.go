package tiermanagement

import (
	"kodegiri/app/model"
	"kodegiri/app/services"

	"github.com/gofiber/fiber/v2"
)

// DeleteTierManagement
// @Summary      Tier Delete function
// @Description  Delete a tier
// @Tags TierManagement
// @Params id path string true "Tier id"
// @Success 200
// @Failure 404
func DeleteTierManagement(c *fiber.Ctx) error {
	db := services.DB

	mod := db.First(&model.Tier{}, "id = ?", c.Params("id"))
	if mod.RowsAffected < 1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "tier not found",
		})
	}

	db.Where("id = ?", c.Params("id")).Delete(&model.Tier{})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully delete tier",
	})
}
