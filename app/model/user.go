package model

import "github.com/go-openapi/strfmt"

type Users struct {
	Base
	UsersAPI
}

type UsersAPI struct {
	Email         *strfmt.Email    `json:"email" gorm:"type:varchar(100), uniqueIndex"`
	Password      *strfmt.Password `json:"-" gorm:"type:longtext"`
	RememberToken *string          `json:"-" gorm:"type:longtext, uniqueIndex"`
	Name          *string          `json:"name" gorm:"type:varchar(100)"`
	PhoneNumber   *string          `json:"phone_number" gorm:"type:varchar(20)"`
	BirthDate     *strfmt.Date     `json:"birth_date" gorm:"type:date"`
	Address       *string          `json:"address" gorm:"type:longtext"`
	JoinDate      *strfmt.Date     `json:"join_date" gorm:"type:date"`
	Status        *string          `json:"status" gorm:"type:varchar(10)"`
}
