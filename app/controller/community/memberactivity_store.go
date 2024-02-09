package community

import (
	"kodegiri/app/model"
	"kodegiri/app/services"
	"time"

	"github.com/gofiber/fiber/v2"
)

// StoreMemberActivity func
// @Summary      Login function
// @Description  Get token for auth
// @Security ApiKeyAuth
// @Tags         Auth
// @Accept       json
// @Produce		 json
// @Param        body body MemberActivityPayload true "Member Activity Payload"
// @Success      200
// @Router       /member-activity [post]
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

	loyaltyprogram := model.LoyaltyProgram{}
	q := db.Where("policy_is_community_activity = ?", true).Or("loyalty_start <= ?", time.Now()).Or("loyalty_end >= ?", time.Now()).Order("created_at DESC").First(&loyaltyprogram)
	if q.RowsAffected > 0 {
		pointTransaction := model.PointTransaction{}
		point := model.Point{}
		db.FirstOrCreate(&point, "user_id = ?", member.ID)

		transType := "earned"
		pointTransaction.Type = &transType
		pointTransaction.Value = loyaltyprogram.BenefitCommunityFixedPoint
		pointTransaction.RemainingPoint = point.Point
		pointTransaction.LoyaltyProgramID = loyaltyprogram.ID
		pointTransaction.TransactionID = mact.TransactionID
		db.Create(&pointTransaction)

		newPoint := *point.Point + *pointTransaction.Value
		point.Point = &newPoint
		db.Save(&point)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Member get member successfully created",
	})
}

type MemberActivityPayload struct {
	MemberID     string `json:"member_id"`
	ActivityName string `json:"activity_name"`
}
