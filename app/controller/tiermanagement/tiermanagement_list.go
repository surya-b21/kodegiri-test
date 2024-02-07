package tiermanagement

import (
	"kodegiri/app/model"
	"kodegiri/app/services"

	"github.com/gofiber/fiber/v2"
)

// ListTierManagement
// @Summary      Tier List function
// @Description  Get all tier list
// @Tags TierManagement
// @Success 200 {object} []model.Tier
// @Failure 404
func ListTierManagement(c *fiber.Ctx) error {
	db := services.DB

	var list []model.Tier
	mod := db.Find(&list)

	if mod.RowsAffected < 1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(list)
}
