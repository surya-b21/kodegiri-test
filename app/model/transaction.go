package model

import (
	"fmt"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Transaction model
type Transaction struct {
	Base
	TransactionAPI
	Users *User
	Item  []*TransactionItem
}

type TransactionAPI struct {
	TotalAmount     *int             `json:"total_amount,omitempty" gorm:"type:integer"`
	TransactionDate *strfmt.DateTime `json:"transaction_date,omitempty" gorm:"type:date"`
	TransactionID   *string          `json:"transaction_id,omitempty" gorm:"type:varchar(50)"`
	UserID          *uuid.UUID       `json:"user_id,omitempty"`
}

type TransactionItem struct {
	Base
	TransactionItemAPI
}

type TransactionItemAPI struct {
	TransactionID *uuid.UUID `json:"transaction_id,omitempty" gorm:"type:varchar(36)"`
	ItemName      *string    `json:"item_name,omitempty" gorm:"type:varchar(50)"`
	ItemPrice     *int       `json:"item_price,omitempty" gorm:"type:bigint"`
	ItemQty       *int       `json:"item_qty,omitempty" gorm:"type:integer"`
	ItemSubtotal  *int       `json:"item_subtotal,omitempty" gorm:"type:integer"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) error {
	count := int64(0)
	now := strfmt.DateTime(time.Now())
	tx.Model(t).Count(&count)

	transactionID := fmt.Sprintf("TRINV/%06d/%s", count+1, time.Now().Format("01022006"))
	t.TransactionID = &transactionID
	t.TransactionDate = &now
	return nil
}
