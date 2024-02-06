package services

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// global variable DB
var DB *gorm.DB

func InitDB() {
	if DB == nil {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d TimeZone=Asia/Jakarta", "postgres", "dockerdb", "S3cret", "postgres", 5432)
		config := &gorm.Config{
			Logger: logger.New(
				log.New(os.Stderr, "[GORM] ", log.LstdFlags), // io writer
				logger.Config{
					SlowThreshold:             time.Second, // Slow SQL threshold
					LogLevel:                  logger.Info, // Log level
					IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
					Colorful:                  true,        // Disable color
				},
			),
			NamingStrategy: schema.NamingStrategy{SingularTable: true},
		}

		if db, err := gorm.Open(postgres.Open(dsn), config); err == nil {
			AutoMigrate(db)
			SeedAll(db)
			DB = db.Debug()
		}
	}
}
