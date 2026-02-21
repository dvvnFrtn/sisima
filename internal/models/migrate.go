// Package model defines data structures used by the data access layer for persistence and retrieval.
package model

import "github.com/dvvnFrtn/sisima/internal/config"

func Migrate() {
	db := config.DB

	if err := CreatePostgresEnums(); err != nil {
		panic(err)
	}

	err := db.AutoMigrate(
		&Student{}, &BillingType{}, &Billing{},
		// &OtherModel{},
	)
	if err != nil {
		panic(err)
	}
}
