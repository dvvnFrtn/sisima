package model

import (
	"gorm.io/gorm"
)

func Migrate() {
	// db := *gorm.DB

	var db *gorm.DB
	if err := CreatePostgresEnums(); err != nil {
		panic(err)
	}

	err := db.AutoMigrate(
		&Student{},
		// &OtherModel{},
	)

	if err != nil {
		panic(err)
	}
}
