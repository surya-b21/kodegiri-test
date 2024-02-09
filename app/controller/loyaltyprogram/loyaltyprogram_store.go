package loyaltyprogram

import (
	"kodegiri/app/model"
	"kodegiri/app/services"

	"github.com/gofiber/fiber/v2"
)

// StoreLoyaltyProgram func
// @Summary      Store Loyalty Program
// @Description  Store a loyalty program
// @Tags LoyaltyProgram
// @Security ApiKeyAuth
// @Param        body body LoyaltyProgramPayload true "Loyalty Program Payload"
// @Success 200
// @Failure 400
// @Router       /loyalty-program [post]
func StoreLoyaltyProgram(c *fiber.Ctx) error {
	payload := LoyaltyProgramPayload{}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Payload",
		})
	}

	db := services.DB
	tx := db.Begin()
	loyaltyTier := []model.LoyaltyProgramTier{}
	loyaltyProgram := model.LoyaltyProgram{}
	loyaltyProgram.Name = payload.Name
	loyaltyProgram.LoyaltyStart = payload.LoyaltyStart
	loyaltyProgram.LoyaltyEnd = payload.LoyaltyEnd
	loyaltyProgram.PolicyTransactionAmount = payload.PolicyTransactionAmount
	loyaltyProgram.PolicyTransactionQty = payload.PolicyTransactionQty
	loyaltyProgram.PolicyIsTransactionFirstPurchase = payload.PolicyIsTransactionFirstPurchase
	loyaltyProgram.PolicyIsCommunityMemberActivity = payload.PolicyIsCommunityMemberActivity
	loyaltyProgram.PolicyIsCommunityMemberGetMember = payload.PolicyIsCommunityMemberGetMember
	loyaltyProgram.PolicyBirthdayBonus = payload.PolicyBirthdayBonus
	loyaltyProgram.BenefitCommunityFixedPoint = payload.BenefitCommunityFixedPoint
	loyaltyProgram.BenefitTransactionFixedPoint = payload.BenefitTransactionFixedPoint
	loyaltyProgram.BenefitTransactionalPrecentage = payload.BenefitTransactionalPrecentage
	tx.Create(&loyaltyProgram)

	for _, tier := range payload.Tier {
		loyaltyTier = append(loyaltyTier, model.LoyaltyProgramTier{
			LoyaltyProgramTierAPI: model.LoyaltyProgramTierAPI{
				LoyaltyProgramID: tier.LoyaltyProgramID,
				TierID:           tier.TierID,
			},
		})
	}
	tx.CreateInBatches(&loyaltyTier, 100)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Loyalty Program successfully created",
	})

}

type LoyaltyProgramPayload struct {
	model.LoyaltyProgramAPI
	Tier []model.LoyaltyProgramTierAPI
}
