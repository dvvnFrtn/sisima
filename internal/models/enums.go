package model

import (
	"gorm.io/gorm"
)

// Create Enums
func CreatePostgresEnums() error {
	// db := config.DB

	var db *gorm.DB
	genderSQL := `
		DO $$ BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender_enum') THEN
			CREATE TYPE gender_enum AS ENUM ('MALE','FEMALE');
		END IF;
		END$$;
	`

	if err := db.Exec(genderSQL).Error; err != nil {
		return err
	}

	return nil
}

// Type Enums
type Gender string

const (
	Male   Gender = "MALE"
	Female Gender = "FEMALE"
)
