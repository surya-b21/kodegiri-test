package community

import (
	"kodegiri/app/model"
	"kodegiri/app/services"

	"github.com/gofiber/fiber/v2"
)

// StoreMemberActivity func
func StoreMemberActivity(c *fiber.Ctx) error {
	payload := MemberActivityPayload{}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Payload",
		})
	}

	db := services.DB

	member := model.User{}
	mod := db.First(&member, "id = ?", payload.MemberID)
	if mod.RowsAffected < 1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Member not found",
		})
	}

	mact := model.MemberActivity{}
	mact.MemberID = member.ID
	mact.ActivityName = &payload.ActivityName
	db.Create(&mact)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Member get member successfully created",
	})
}

type MemberActivityPayload struct {
	MemberID     string `json:"member_id"`
	ActivityName string `json:"activity_name"`
}
