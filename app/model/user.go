package model

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-openapi/strfmt"
	"golang.org/x/crypto/bcrypt"
)

// Users Model
type Users struct {
	Base
	UsersAPI
	ReedemedPoint    int `json:"reedemed_point,omitempty"`
	EarnedPoint      int `json:"earned_point,omitempty"`
	RemaiendPoint    int `json:"remained_point,omitempty"`
	Point            *Point
	PointTransaction []*PointTransaction
}

type UsersAPI struct {
	Email         *strfmt.Email    `json:"email,omitempty" gorm:"type:varchar(100); uniqueIndex"`
	Password      *strfmt.Password `json:"-" gorm:"type:text"`
	RememberToken *string          `json:"-" gorm:"type:text; uniqueIndex"`
	Name          *string          `json:"name,omitempty" gorm:"type:varchar(100)"`
	PhoneNumber   *string          `json:"phone_number,omitempty" gorm:"type:varchar(20)"`
	BirthDate     *strfmt.Date     `json:"birth_date,omitempty" gorm:"type:date"`
	Address       *string          `json:"address,omitempty" gorm:"type:text"`
	JoinDate      *strfmt.Date     `json:"join_date,omitempty" gorm:"type:date"`
	Status        *string          `json:"status,omitempty" gorm:"type:varchar(10)"`
	MemberNumber  *string          `json:"member_number,omitempty" gorm:"type:varchar(5);uniqueIndex"`
}

func (s *Users) Seed() *[]Users {
	hash, _ := bcrypt.GenerateFromPassword([]byte("$Password123"), bcrypt.MinCost)
	password := strfmt.Password(string(hash))
	email := "test@mail.com"
	name := "Test Account"
	birthDate := strfmt.Date(time.Now().AddDate(-20, 0, 0))
	joinDate := strfmt.Date(time.Now().AddDate(0, -3, 0))
	status := "active"
	memberNumber := fmt.Sprintf("%05d", 10000+rand.Intn(90000))

	users := []Users{
		{
			UsersAPI: UsersAPI{
				Email:        (*strfmt.Email)(&email),
				Password:     &password,
				Name:         &name,
				BirthDate:    &birthDate,
				JoinDate:     &joinDate,
				Status:       &status,
				MemberNumber: &memberNumber,
			},
		},
	}

	return &users
}
