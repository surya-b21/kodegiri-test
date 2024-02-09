package model

type Tier struct {
	Base
	TierAPI
}

type TierAPI struct {
	Name         *string `json:"name,omitempty" gorm:"type:varchar(100)"`
	MinimalPoint *int    `json:"minimal_point,omitempty" gorm:"type:integer"`
	MaximalPoint *int    `json:"maximal_point,omitempty" gorm:"type:integer"`
}
