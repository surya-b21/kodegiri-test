package membership

import (
	"kodegiri/app/model"
	"kodegiri/app/services"

	"github.com/gofiber/fiber/v2"
)

// Membership Detail
// @Summary      Membership Detail function
// @Description  Get membership detail
// @Params id path string false "membership id"
// @Tags Membership
// @Success 200
// @Failure 404
func DetailMembership(c *fiber.Ctx) error {
	db := services.DB

	membership := model.Users{}
	mod := db.Joins("Point").Joins("PointTransaction").First(&membership, `id = ?`, c.Params("id"))
	if mod.RowsAffected < 1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Membership not found",
		})
	}

	membership.RemaiendPoint = *membership.Point.Point

	for _, transaction := range membership.PointTransaction {
		transaction.TransactionDate = *transaction.CreatedAt
		switch *transaction.Type {
		case "earned":
			membership.EarnedPoint += *transaction.Value
		case "redeemed":
			membership.ReedemedPoint += *transaction.Value
		}
	}

	return c.Status(fiber.StatusOK).JSON(membership)
}
