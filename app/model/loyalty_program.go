package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

// LoyaltyProgram model
type LoyaltyProgram struct {
	Base
	LoyaltyProgramAPI
	LoyaltyProgramTier []LoyaltyProgramTier
}

type LoyaltyProgramAPI struct {
	Name                             *string          `json:"name,omitempty" gorm:"type:varchar(50)"`
	LoyaltyStart                     *strfmt.DateTime `json:"loyalty_start,omitempty" gorm:"type:datetime" swaggertype:"string"`
	LoyaltyEnd                       *strfmt.DateTime `json:"loyalty_end,omitempty" gorm:"type:datetime" swaggertype:"string"`
	PolicyTransactionAmount          *int             `json:"policy_transaction_amount,omitempty" gorm:"type:integer"`
	PolicyTransactionQty             *int             `json:"policy_transaction_qty,omitempty" gorm:"type:integer"`
	PolicyIsTransactionFirstPurchase *bool            `json:"policy_is_transaction_first_purchase,omitempty" gorm:"type:bool"`
	PolicyIsCommunityMemberGetMember *bool            `json:"policy_is_community_member_get_member,omitempty" gorm:"type:bool"`
	PolicyIsCommunityMemberActivity  *bool            `json:"policy_is_community_activity,omitempty" gorm:"type:bool"`
	PolicyBirthdayBonus              *bool            `json:"policy_birthday_bonus,omitempty" gorm:"type:bool"`
	BenefitTransactionalPrecentage   *int             `json:"benefit_transaction_precentage,omitempty" gorm:"type:bool"`
	BenefitTransactionFixedPoint     *int             `json:"benefit_transaction_fixed_point,omitempty" gorm:"type:integer"`
	BenefitCommunityFixedPoint       *int             `json:"benefit_community_fixed_point" gorm:"type:integer"`
}

type LoyaltyProgramTier struct {
	Base
	LoyaltyProgramTierAPI
	LoyaltyProgram *LoyaltyProgram
	Tier           *Tier
}

type LoyaltyProgramTierAPI struct {
	LoyaltyProgramID *uuid.UUID `json:"loyalty_program_id" gorm:"type:varchar(36)"`
	TierID           *uuid.UUID `json:"tier_id" gorm:"type:varchar(36)"`
}
