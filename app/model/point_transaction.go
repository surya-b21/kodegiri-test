package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

// PoinTransaction model
type PointTransaction struct {
	Base
	PointTransactionAPI
	TransactionDate strfmt.DateTime `json:"transaction_date,omitempty"`
}

type PointTransactionAPI struct {
	Type    *string    `json:"type,omitempty" gorm:"type:varchar(10)"`
	Value   *int       `json:"value,omitempty" gorm:"type:integer"`
	UserID  *uuid.UUID `json:"user_id,omitempty" gorm:"type:varchar(36)"`
	PointID *uuid.UUID `json:"point_id,omitempty" gorm:"type:varchar(36)"`
}
