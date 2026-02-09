package model

import "github.com/dvvnFrtn/sisima/internal/config"

func Migrate() {
	db := config.DB

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
