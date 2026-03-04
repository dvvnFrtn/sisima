package dtodata

import (
	model "github.com/dvvnFrtn/sisima/internal/models"
	"github.com/google/uuid"
)

type CreateBillingTypeRequest struct {
	Name      string               `json:"name" validate:"required"`
	Amount    int64                `json:"amount" validate:"required,gt=0"`
	Recurring billingTypeRecurring `json:"recurring" validate:"required"`
}

type UpdateBillingTypeRequest struct {
	Name   *string `json:"name,omitempty"`
	Amount *int64  `json:"amount,omitempty" validate:"gt=0"`
}

type BillingTypeResponse struct {
	ID        uuid.UUID            `json:"id"`
	Name      string               `json:"name"`
	Amount    int64                `json:"amount"`
	Recurring billingTypeRecurring `json:"recurring"`
}

type billingTypeRecurring struct {
	Interval      string `json:"interval" validate:"required,oneof=MONTH YEAR ONCE"`
	IntervalCount *int64 `json:"interval_count" validate:"required"`
}

func ToBillingTypeResponse(bt model.BillingType) BillingTypeResponse {
	return BillingTypeResponse{
		ID:     bt.ID,
		Name:   bt.Name,
		Amount: bt.Amount,
		Recurring: billingTypeRecurring{
			Interval:      string(bt.Interval),
			IntervalCount: &bt.IntervalCount,
		},
	}
}

type CreateBillingRequest struct {
	StudentID     string `json:"student_id" validate:"required"`
	BillingTypeID string `json:"billing_type_id" validate:"required"`
	Period        string `json:"period" validate:"required"`
}

type BillingResponse struct {
	ID            uuid.UUID `json:"id"`
	StudentID     uuid.UUID `json:"student_id"`
	BillingTypeID uuid.UUID `json:"billing_type_id"`
	Period        string    `json:"period"`
	Amount        int64     `json:"amount"`
	Status        string    `json:"status"`
}

func ToBillingResponse(b model.Billing) BillingResponse {
	return BillingResponse{
		ID:            b.ID,
		StudentID:     b.StudentID,
		BillingTypeID: b.BillingTypeID,
		Amount:        b.Amount,
		Period:        b.Period.Format("02-01-2006"),
		Status:        string(b.Status),
	}
}

func Map[T any, U any](data []T, f func(T) U) []U {
	result := make([]U, len(data))
	for i, v := range data {
		result[i] = f(v)
	}

	return result
}
