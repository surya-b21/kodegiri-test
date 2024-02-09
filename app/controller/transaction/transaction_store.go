package transaction

import (
	"kodegiri/app/model"
	"kodegiri/app/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// StoreTransaction
// @Summary      StoreTransaction function
// @Description  Store a transaction
// @Security ApiKeyAuth
// @Tags StoreTransaction
// @Params body json TransactionPayload true "Transaction payload"
// @Success 200
// @Failure 400
// @Router       /transaction [post]
func StoreTransaction(c *fiber.Ctx) error {
	payload := TransactionPayload{}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Payload",
		})
	}

	db := services.DB
	transaction := model.Transaction{}
	transactionItems := []model.TransactionItem{}
	user := model.User{}
	mod := db.Find(&user, "id = ?", payload.UserID)
	if mod.RowsAffected < 1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	tx := db.Begin()
	parsedID, err := uuid.Parse(payload.UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user id",
		})
	}

	firstPurchase := false
	mod = tx.Model(&model.Transaction{}).Where("user_id = ?", user.ID).First(&model.Transaction{})
	if mod.RowsAffected < 1 {
		firstPurchase = true
	}

	totalQty := 0
	for _, item := range payload.Item {
		totalQty += *item.ItemQty
	}

	// find loyalty policy
	loyaltyProgram := model.LoyaltyProgram{}
	q := tx.Where("policy_is_transaction_first_purchase = ?", firstPurchase).
		Or("policy_transaction_amount <= ?", payload.TotalAmount).
		Or("policy_transaction_qty <= ?", totalQty).
		Or("loyalty_start <= ?", time.Now()).
		Or("loyalty_end >= ?", time.Now()).
		Order("created_at DESC").
		First(&loyaltyProgram)
	if q.RowsAffected > 0 {
		pointTransaction := model.PointTransaction{}
		point := model.Point{}
		tx.FirstOrCreate(&point, "user_id = ?", user.ID)

		transType := "earned"
		pointTransaction.Type = &transType
		pointTransaction.Value = loyaltyProgram.BenefitCommunityFixedPoint
		pointTransaction.RemainingPoint = point.Point
		pointTransaction.LoyaltyProgramID = loyaltyProgram.ID
		tx.Create(&pointTransaction)

		newPoint := *point.Point + *pointTransaction.Value
		point.Point = &newPoint
		tx.Save(&point)
	}

	transaction.UserID = &parsedID
	transaction.TotalAmount = &payload.TotalAmount
	tx.Create(&transaction)

	for _, item := range payload.Item {
		transactionItems = append(transactionItems, model.TransactionItem{
			TransactionItemAPI: model.TransactionItemAPI{
				TransactionID: transaction.ID,
				ItemName:      item.ItemName,
				ItemPrice:     item.ItemPrice,
				ItemQty:       item.ItemQty,
				ItemSubtotal:  item.ItemSubtotal,
			},
		})
	}
	tx.CreateInBatches(transactionItems, 100)

	tx.Commit()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Successfully create transaction",
	})
}

type TransactionPayload struct {
	UserID      string `json:"user_id"`
	Item        []model.TransactionItemAPI
	TotalAmount int `json:"total_amount,omitempty"`
}
