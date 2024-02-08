package redeemedpoint

import (
	"kodegiri/app/model"
	"kodegiri/app/services"

	"github.com/gofiber/fiber/v2"
)

// ReedemedPointStore func
// @Summary      ReedemedPointStore function
// @Description  Store a reedemed point
// @Tags ReedemedPointStore
// @Params body json ReedemedPointPayload true "ReedemedPoint payload"
// @Success 200
// @Failure 400
func ReedemedPointStore(c *fiber.Ctx) error {
	payload := ReedemedPointPayload{}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Payload",
		})
	}

	db := services.DB
	membership := model.User{}
	mod := db.Joins("Point").First(&membership, `id = ?`, c.Params("id"))
	if mod.RowsAffected < 1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Membership not found",
		})
	}

	// validate redeeem and remaining
	if payload.Type == "reedemed" && *membership.Point.Point <= payload.Point {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Reedemed point can't greater than remaining point",
		})
	}

	remainingPoint := *membership.Point.Point
	if payload.Type == "earned" {
		remainingPoint += payload.Point
	} else {
		remainingPoint -= payload.Point
	}

	tx := db.Begin()
	history := model.PointTransaction{}
	history.Type = &payload.Type
	history.UserID = membership.ID
	history.RemainingPoint = &remainingPoint
	history.PointID = membership.Point.ID
	tx.Create(&history)

	tx.Where("id = ?", history.PointID).Update("point", remainingPoint)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully reedemed point",
	})

}

type ReedemedPointPayload struct {
	MemberID string `json:"member_id"`
	Type     string `json:"type"`
	Point    int    `json:"redeemed_point"`
}
