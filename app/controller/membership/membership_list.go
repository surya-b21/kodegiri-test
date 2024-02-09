package membership

import (
	"kodegiri/app/model"
	"kodegiri/app/services"

	"github.com/gofiber/fiber/v2"
)

// Membership
// @Summary      Membership List function
// @Description  Get all membership list
// @Security ApiKeyAuth
// @Tags Membership
// @Success 200
// @Failure 404
// @Router       /membership [get]
func ListMembership(c *fiber.Ctx) error {
	db := services.DB

	list := []model.User{}
	mod := db.Joins("Point").Find(&list)

	if mod.RowsAffected < 1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(list)
}
