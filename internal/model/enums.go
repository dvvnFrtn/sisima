package model

import "github.com/dvvnFrtn/sisima/internal/config"

// Create Enums
func CreatePostgresEnums() error {
	db := config.DB

	genderSQL := `
		DO $$ BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender_enum') THEN
			CREATE TYPE gender_enum AS ENUM ('MALE','FEMALE');
		END IF;
		END$$;
	`

	classSQL := `
		DO $$ BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'class_enum') THEN
			CREATE TYPE class_enum AS ENUM ('0','1','2','3','4','5','6','7');
		END IF;
		END$$;
	`

	if err := db.Exec(genderSQL).Error; err != nil {
		return err
	}
	if err := db.Exec(classSQL).Error; err != nil {
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

type Class int

const (
	Class0 Class = 0
	Class1 Class = 1
	Class2 Class = 2
	Class3 Class = 3
	Class4 Class = 4
	Class5 Class = 5
	Class6 Class = 6
	Class7 Class = 7
)
