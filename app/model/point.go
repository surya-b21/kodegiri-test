package model

import (
	"crypto/rand"
	"encoding/base64"
	"errors"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Poin Model
type Point struct {
	Base
	PointAPI
	PointTransaction []*PointTransaction
}

type PointAPI struct {
	Referral *string    `json:"referral,omitempty" gorm:"type:varchar(7);uniqueIndex"`
	Point    *int       `json:"point,omitempty" gorm:"type:integer"`
	UserID   *uuid.UUID `json:"user_id,omitempty" gorm:"type:varchar(36)"`
}

func (s *Point) BeforeCreate(tx *gorm.DB) error {
	b := make([]byte, 7)
	if _, err := rand.Read(b); err != nil {
		return errors.New("Failed generate referral")
	}
	referral := base64.URLEncoding.EncodeToString(b)
	point := 0
	s.Referral = &referral
	s.Point = &point
	return nil
}

type PointTransaction struct {
	Base
	PointTransactionAPI
	TransactionDate strfmt.DateTime `json:"transaction_date,omitempty"`
	User            *User
	LoyaltyProgram  *LoyaltyProgram
	MemberGetMember *MemberGetMember `json:"member_get_member,omitempty" gorm:"foreignKey:TransactionID; references:TransactionID"`
	MemberActivity  *MemberActivity  `json:"member_activity,omitempty" gorm:"foreignKey:TransactionID; references:TransactionID"`
	Transaction     *Transaction     `json:"transaction,omitempty" gorm:"foreignKey:TransactionID; references:TransactionID"`
}

type PointTransactionAPI struct {
	Type             *string    `json:"type,omitempty" gorm:"type:varchar(10)"`
	Value            *int       `json:"value,omitempty" gorm:"type:integer"`
	RemainingPoint   *int       `json:"remaining_point,omitempty" gorm:"type:integer"`
	UserID           *uuid.UUID `json:"user_id,omitempty" gorm:"type:varchar(36)"`
	PointID          *uuid.UUID `json:"point_id,omitempty" gorm:"type:varchar(36)"`
	LoyaltyProgramID *uuid.UUID `json:"loyalty_program_id,omitempty" gorm:"type:varchar(36)"`
	TransactionID    *string    `json:"transaction_id,omitempty" gorm:"type:varchar(100);uniqueIndex"`
}
