package services

import (
	"errors"
	"fmt"
	"kodegiri/app/model"
	"os"
	"reflect"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedAll(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		seeds := DataSeeds()
		for i := range seeds {
			err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(seeds[i]).Error
			if os.Getenv("SIMULATE_SEEDER_FAIL") != "" {
				err = errors.New(os.Getenv("SIMULATE_SEEDER_FAIL"))
			}
			if nil != err {
				name := reflect.TypeOf(seeds[i]).String()
				errorMessage := err.Error()
				return fmt.Errorf("%s seeder fail with %s", name, errorMessage)
			}
		}
		return nil
	})
}

var (
	user model.Users
)

// DataSeeds data to seeds
func DataSeeds() []interface{} {
	return []interface{}{
		user.Seed(),
	}
}
