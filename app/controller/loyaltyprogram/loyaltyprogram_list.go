package loyaltyprogram

import (
	"kodegiri/app/model"
	"kodegiri/app/services"

	"github.com/gofiber/fiber/v2"
)

// ListLoyaltyProgram func
// @Summary      LoyaltiProgramList function
// @Description  Get all loyalty program list
// @Tags LoyaltyProgram
// @Security ApiKeyAuth
// @Success 200 {object} []model.LoyaltyProgram
// @Failure 404
// @Router       /loyalty-program [get]
func ListLoyaltyProgram(c *fiber.Ctx) error {
	db := services.DB

	list := []model.LoyaltyProgram{}
	mod := db.Joins(`LoyaltyProgramTier.Tier`).Find(&list)
	if mod.RowsAffected < 1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&list)
}
