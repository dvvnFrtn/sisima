package dto

import (
	"time"

	model "github.com/dvvnFrtn/sisima/internal/models"
	"github.com/google/uuid"
)

// Exception
type ExceptionTitle string

const (
	InvalidRequest ExceptionTitle = "INVALID_REQUEST"
	ValidationErr  ExceptionTitle = "VALIDATION_ERROR"
	InternalErr    ExceptionTitle = "INTERNAL_ERROR"
)

type ExceptionResponse struct {
	Title  ExceptionTitle `json:"title"`
	Errors interface{}    `json:"errors,omitempty"`
}

func NewExceptionResponse(title ExceptionTitle, errors interface{}) ExceptionResponse {
	return ExceptionResponse{
		Title:  title,
		Errors: errors,
	}
}

// Other
type CreateStudentRequest struct {
	ID        uuid.UUID `json:"id"`
	NIS       string    `json:"nis" validate:"required,min=1,max=22"`
	NISN      string    `json:"nisn" validate:"required,min=1,max=22"`
	FullName  string    `json:"full_name" validate:"required,min=1,max=70"`
	NickName  string    `json:"nick_name" validate:"omitempty,min=1,max=20"`
	Gender    string    `json:"gender" validate:"required,oneof=MALE FEMALE"`
	EntryYear string    `json:"entry_year" validate:"required,len=4,numeric"`
	Class     string    `json:"class" validate:"required,len=1,alphaunicode"`
}

type StudentResponse struct {
	ID        uuid.UUID `json:"id"`
	NIS       string    `json:"nis"`
	NISN      string    `json:"nisn"`
	FullName  string    `json:"full_name"`
	NickName  string    `json:"nick_name"`
	Gender    string    `json:"gender"`
	EntryYear string    `json:"entry_year"`
	Class     string    `json:"class"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToStudentResponse(s *model.Student) *StudentResponse {
	if s == nil {
		return nil
	}
	return &StudentResponse{
		ID:        s.ID,
		NIS:       s.NIS,
		NISN:      s.NISN,
		FullName:  s.FullName,
		NickName:  s.NickName,
		Gender:    string(s.Gender),
		EntryYear: s.EntryYear,
		Class:     s.Class,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

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
