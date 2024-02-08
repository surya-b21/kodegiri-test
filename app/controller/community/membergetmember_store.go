package community

import (
	"kodegiri/app/model"
	"kodegiri/app/services"

	"github.com/go-openapi/strfmt"
	"github.com/gofiber/fiber/v2"
)

// StoreMemberGetMember func
// @Summary     StoreMemberGetMember function
// @Description  Store a Member Get Member
// @Tags Community
// @Params body json MemberGetMemberPayload true "Tier member get member"
// @Success 200
// @Failure 400
// @Failure 404
func StoreMemberGetMember(c *fiber.Ctx) error {
	payload := MemberGetMemberPayload{}
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
	mgm := model.MemberGetMember{}
	mgm.MemberID = member.ID
	mgm.PersonEmail = (*strfmt.Email)(&payload.PersonEmail)
	mgm.PersonName = &payload.PersonName
	mgm.PersonPhoneNumber = &payload.PersonPhoneNumber
	db.Create(&mgm)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Member get member successfully created",
	})
}

type MemberGetMemberPayload struct {
	MemberID          string `json:"member_id"`
	PersonName        string `json:"person_name"`
	PersonPhoneNumber string `json:"person_phone_number"`
	PersonEmail       string `json:"person_email"`
}
