// Package service implements business logic and application use cases.
package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/dvvnFrtn/sisima/internal/config"
	model "github.com/dvvnFrtn/sisima/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BillingService interface {
	CreateBillingType(param CreateBillingTypeRequest) error
}

type billingService struct{}

func NewBillingService() BillingService {
	return &billingService{}
}

type CreateBillingTypeRequest struct {
	Name          string
	DefaultAmount int64
	Recurring     *struct {
		Interval      string
		IntervalCount int64
	}
}

func (bs *billingService) CreateBillingType(param CreateBillingTypeRequest) error {
	bt := model.BillingType{
		ID:     uuid.New(),
		Name:   param.Name,
		Amount: param.DefaultAmount,
	}

	if param.Recurring != nil {
		if !model.BillingTypeInterval(param.Recurring.Interval).IsValid() {
			return fmt.Errorf("invalid interval value '%s'", param.Recurring.Interval)
		}
		bt.Interval = model.BillingTypeInterval(param.Recurring.Interval)

		if param.Recurring.Interval != string(model.BillingTypeIntevalOnce) && param.Recurring.IntervalCount == 0 {
			return fmt.Errorf("interval_count value can't be zero for interval '%s'", param.Recurring.Interval)
		}
		bt.IntervalCount = param.Recurring.IntervalCount
	}

	if err := config.DB.Create(&bt).Error; err != nil {
		return fmt.Errorf("insert billing_type: %w: %T", err, err)
	}

	return nil
}

func (bs *billingService) GetAllBillingType() error {
	var bts []model.BillingType
	if err := config.DB.Find(&bts).Error; err != nil {
		return fmt.Errorf("find billing_type: %w", err)
	}

	return nil
}

func (bs *billingService) GetBillingType(ID uuid.UUID) error {
	var bt model.BillingType
	if err := config.DB.Find(&bt, ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("billing_type with id '%s' could'nt be found", ID.String())
		}
		return fmt.Errorf("find billing_type: %w", err)
	}

	return nil
}

type UpdateBillingTypeRequest struct {
	ID            uuid.UUID
	Name          string
	DefaultAmount int64
}

func (bs *billingService) UpdateBillingType(param UpdateBillingTypeRequest) error {
	var bt model.BillingType
	if err := config.DB.Find(&bt, param.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("billing_type with id '%s' could'nt be found", param.ID)
		}
		return fmt.Errorf("find billing_type: %w", err)
	}

	bt.Name = param.Name
	bt.Amount = param.DefaultAmount

	if err := config.DB.Save(&bt).Error; err != nil {
		return fmt.Errorf("save billing_type: %w", err)
	}

	return nil
}

type CreateBillingRequest struct {
	StudentID     uuid.UUID
	BillingTypeID uuid.UUID
	Period        time.Time
}

func (bs *billingService) CreateBilling(param CreateBillingRequest) error {
	var bt model.BillingType
	if err := config.DB.Find(&bt, param.BillingTypeID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("billing_type with id '%s' could'nt be found", param.BillingTypeID)
		}
		return fmt.Errorf("find billing_type: %w", err)
	}

	var period time.Time
	switch bt.Interval {
	case model.BillingTypeIntervalYear:
		period = bs.truncateToYear(param.Period)
	case model.BillingTypeIntervalMonth:
		period = bs.truncateToMonth(param.Period)
	default:
		period = param.Period
	}

	b := model.Billing{
		ID:            uuid.New(),
		BillingTypeID: bt.ID,
		StudentID:     param.StudentID,
		Amount:        bt.Amount,
		Interval:      bt.Interval,
		Status:        model.BillingStatusUnpaid,
		Period:        period,
	}
	if err := config.DB.Create(&b).Error; err != nil {
		return fmt.Errorf("create billing: %w", err)
	}

	return nil
}

type UpdateBillingRequest struct {
	ID     uuid.UUID
	Amount int64
	Status string
}

func (bs *billingService) UpdateBilling(param UpdateBillingRequest) error {
	var bl model.Billing
	if err := config.DB.First(&bl, param.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("billing with id '%s' could'nt be found", param.ID)
		}
		return fmt.Errorf("find billing: %w", err)
	}

	bl.Amount = param.Amount
	bl.Status = model.BillingStatus(param.Status)

	if err := config.DB.Save(&bl).Error; err != nil {
		return fmt.Errorf("save billing: %w", err)
	}

	return nil
}

func (bs *billingService) GetAllBilling() error {
	var bls []model.Billing
	if err := config.DB.Find(&bls).Error; err != nil {
		return fmt.Errorf("find billing: %w", err)
	}

	return nil
}

func (bs *billingService) GetBilling(ID uuid.UUID) error {
	var bl model.Billing
	if err := config.DB.First(&bl, ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("billing with id '%s' could'nt be found", ID)
		}
		return fmt.Errorf("find billing: %w", err)
	}

	return nil
}

func (bs *billingService) GetStudentBilling(sID uuid.UUID) error {
	var bls []model.Billing
	if err := config.DB.Where("student_id = ?", sID).Find(&bls).Error; err != nil {
		return fmt.Errorf("find billing: %w", err)
	}

	return nil
}

func (bs *billingService) truncateToMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func (bs *billingService) truncateToYear(t time.Time) time.Time {
	return time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, t.Location())
}
