package model

import (
	"fmt"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CommunityMember model
type MemberGetMember struct {
	Base
	MemberGetMemberAPI
	Member *User `json:"member,omitempty" gorm:"foreignKey:MemberID"`
}

type MemberGetMemberAPI struct {
	MemberID          *uuid.UUID       `json:"member_id,omitempty" gorm:"type:varchar(36)"`
	PersonName        *string          `json:"person_name,omitempty" gorm:"type:varchar(50)"`
	PersonPhoneNumber *string          `json:"person_phone_number,omitempty" gorm:"type:varchar(20)"`
	PersonEmail       *strfmt.Email    `json:"person_email,omitempty" gorm:"type:varchar(100);"`
	TransactionDate   *strfmt.DateTime `json:"transaction_date,omitempty" gorm:"type:timestamp" format:"date-time" swaggertype:"string"`
	TransactionID     *string          `json:"transaction_id,omitempty" gorm:"type:varchar(50);uniqueIndex"`
}

func (m *MemberGetMember) BeforeCreate(tx *gorm.DB) error {
	count := int64(0)
	now := strfmt.DateTime(time.Now())
	tx.Model(m).Count(&count)

	transactionID := fmt.Sprintf("TRMGM/%06d/%s", count+1, time.Now().Format("01022006"))
	m.TransactionID = &transactionID
	m.TransactionDate = &now
	return nil
}

// CommunityMemberActivity model
type MemberActivity struct {
	Base
	MemberActivityAPI
	Member *User `json:"member,omitempty" gorm:"foreignKey:MemberID"`
}

type MemberActivityAPI struct {
	ActivityName    *string          `json:"activity_name,omitempty" gorm:"type:varchar(255)"`
	MemberID        *uuid.UUID       `json:"member_id,omitempty" gorm:"type:varchar(36)"`
	TransactionDate *strfmt.DateTime `json:"transaction_date,omitempty" gorm:"type:timestamp" format:"date-time" swaggertype:"string"`
	TransactionID   *string          `json:"transaction_id,omitempty" gorm:"type:varchar(50);uniqueIndex"`
}

func (a *MemberActivity) BeforeCreate(tx *gorm.DB) error {
	count := int64(0)
	now := strfmt.DateTime(time.Now())
	tx.Model(a).Count(&count)

	transactionID := fmt.Sprintf("TRACT/%06d/%s", count+1, time.Now().Format("01022006"))
	a.TransactionID = &transactionID
	a.TransactionDate = &now
	return nil
}
