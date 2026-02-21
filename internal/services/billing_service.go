// Package service implements business logic and application use cases.
package service

import (
	"time"

	"github.com/dvvnFrtn/sisima/internal/config"
	"github.com/dvvnFrtn/sisima/internal/dto"
	model "github.com/dvvnFrtn/sisima/internal/models"
	"github.com/google/uuid"
)

type BillingService interface {
	CreateBillingType(param dto.CreateBillingTypeRequest) error
	UpdateBillingType(tID uuid.UUID, param dto.UpdateBillingTypeRequest) error
	GetAllBillingType() ([]dto.BillingTypeResponse, error)
	GetBillingType(ID uuid.UUID) (dto.BillingTypeResponse, error)

	CreateBilling(param dto.CreateBillingRequest) error
	GetAllBilling() ([]dto.BillingResponse, error)
	GetBilling(ID uuid.UUID) (dto.BillingResponse, error)
}

type billingService struct{}

func NewBillingService() BillingService {
	return &billingService{}
}

func (bs *billingService) CreateBillingType(param dto.CreateBillingTypeRequest) error {
	if err := config.DB.Create(&model.BillingType{
		ID:            uuid.New(),
		Name:          param.Name,
		Amount:        param.Amount,
		Interval:      model.BillingTypeInterval(param.Recurring.Interval),
		IntervalCount: *param.Recurring.IntervalCount,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (bs *billingService) GetAllBillingType() ([]dto.BillingTypeResponse, error) {
	var bts []model.BillingType
	if err := config.DB.Find(&bts).Error; err != nil {
		return []dto.BillingTypeResponse{}, err
	}

	return dto.Map(bts, dto.ToBillingTypeResponse), nil
}

func (bs *billingService) GetBillingType(ID uuid.UUID) (dto.BillingTypeResponse, error) {
	var bt model.BillingType
	if err := config.DB.First(&bt, ID).Error; err != nil {
		return dto.BillingTypeResponse{}, err
	}

	return dto.ToBillingTypeResponse(bt), nil
}

func (bs *billingService) UpdateBillingType(tID uuid.UUID, param dto.UpdateBillingTypeRequest) error {
	var bt model.BillingType
	if err := config.DB.Find(&bt, tID).Error; err != nil {
		return err
	}

	if param.Name != nil {
		bt.Name = *param.Name
	}

	if param.Amount != nil {
		bt.Amount = *param.Amount
	}

	if err := config.DB.Save(&bt).Error; err != nil {
		return err
	}

	return nil
}

func (bs *billingService) CreateBilling(param dto.CreateBillingRequest) error {
	var bt model.BillingType
	if err := config.DB.Find(&bt, uuid.MustParse(param.BillingTypeID)).
		Error; err != nil {
		return err
	}

	periodRaw, err := time.Parse("02-01-2006", param.Period)
	if err != nil {
		return err
	}

	var period time.Time
	switch bt.Interval {
	case model.BillingTypeIntervalYear:
		period = bs.truncateToYear(periodRaw)
	case model.BillingTypeIntervalMonth:
		period = bs.truncateToMonth(periodRaw)
	default:
		period = periodRaw
	}

	if err := config.DB.Create(model.Billing{
		ID:            uuid.New(),
		BillingTypeID: bt.ID,
		StudentID:     uuid.MustParse(param.StudentID),
		Amount:        bt.Amount,
		Interval:      bt.Interval,
		Status:        model.BillingStatusUnpaid,
		Period:        period,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (bs *billingService) GetAllBilling() ([]dto.BillingResponse, error) {
	var bls []model.Billing
	if err := config.DB.Find(&bls).Error; err != nil {
		return []dto.BillingResponse{}, err
	}

	return dto.Map(bls, dto.ToBillingResponse), nil
}

func (bs *billingService) GetBilling(ID uuid.UUID) (dto.BillingResponse, error) {
	var bl model.Billing
	if err := config.DB.First(&bl, ID).Error; err != nil {
		return dto.BillingResponse{}, err
	}

	return dto.ToBillingResponse(bl), nil
}

func (bs *billingService) truncateToMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func (bs *billingService) truncateToYear(t time.Time) time.Time {
	return time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, t.Location())
}
