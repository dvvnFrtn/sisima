package model

import (
	"time"

	"github.com/google/uuid"
)

type BillingType struct {
	ID            uuid.UUID           `gorm:"type:uuid;primaryKey;not null"`
	Name          string              `gorm:"type:varchar;not null"`
	Amount        int64               `gorm:"type:integer;not null"`
	Interval      BillingTypeInterval `gorm:"type:billing_type_interval;not null"`
	IntervalCount int64               `gorm:"type:integer;not null"`
	CreatedAt     time.Time           `gorm:"autoCreateTime;not null"`
	UpdatedAt     time.Time           `gorm:"autoUpdateTime;not null"`
}

type BillingTypeInterval string

const (
	BillingTypeIntervalMonth BillingTypeInterval = "MONTH"
	BillingTypeIntervalYear  BillingTypeInterval = "YEAR"
	BillingTypeIntevalOnce   BillingTypeInterval = "ONCE"
)

func (bti BillingTypeInterval) IsValid() bool {
	switch bti {
	case BillingTypeIntervalMonth, BillingTypeIntervalYear, BillingTypeIntevalOnce:
		return true
	}
	return false
}

type Billing struct {
	ID            uuid.UUID           `gorm:"type:uuid;primaryKey;not null"`
	StudentID     uuid.UUID           `gorm:"type:uuid;not null"`
	BillingTypeID uuid.UUID           `gorm:"type:uuid"`
	Period        time.Time           `gorm:"type:date;not null"`
	Amount        int64               `gorm:"type:integer;not null"`
	Interval      BillingTypeInterval `gorm:"type:billing_type_interval;not null"`
	Status        BillingStatus       `gorm:"type:billing_status;not null"`
	CreatedAt     time.Time           `gorm:"autoCreateTime;not null"`
	UpdatedAt     time.Time           `gorm:"autoUpdateTime;not null"`

	Student     Student
	BillingType BillingType
}

type BillingStatus string

const (
	BillingStatusUnpaid BillingStatus = "UNPAID"
	BillingStatusPaid   BillingStatus = "PAID"
)
