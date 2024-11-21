package migrations

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/go-gormigrate/gormigrate/v2"
)

func addFileSyncers() *gormigrate.Migration {
	type FileSyncer struct {
		Model
		Source       string `gorm:"index"`
		ConsumerName string `gorm:"index"`
		Version      int    `gorm:"not null"`
		// Spec is file syncer spec with JSON format.
		Spec datatypes.JSON `gorm:"type:json"`
		// Status is file syncer status with JSON format.
		Status datatypes.JSON `gorm:"type:json"`
	}

	return &gormigrate.Migration{
		ID: "202411151632",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&FileSyncer{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&FileSyncer{})
		},
	}
}
