package transaction

import (
	"kodegiri/app/model"
	"kodegiri/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// StoreTransaction
// @Summary      StoreTransaction function
// @Description  Store a transaction
// @Tags StoreTransaction
// @Params body json TransactionPayload true "Transaction payload"
// @Success 200
// @Failure 400
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

	tx := db.Begin()
	parsedID, err := uuid.Parse(payload.UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user id",
		})
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
