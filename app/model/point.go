package model

import (
	"crypto/rand"
	"encoding/base64"
	"errors"

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
