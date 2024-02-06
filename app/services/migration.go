package services

import (
	"kodegiri/app/model"

	"gorm.io/gorm"
)

// AutoMigrate function
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(ModelList...)
}

// ModelList variable
var ModelList []interface{} = []interface{}{
	&model.Users{},
}
