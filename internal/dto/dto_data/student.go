package dtodata

import (
	"time"

	model "github.com/dvvnFrtn/sisima/internal/models"
	"github.com/google/uuid"
)

//
// Other
//

// Students
type CreateStudentRequest struct {
	ID        uuid.UUID `json:"id"`
	NIS       string    `json:"nis" validate:"max=22"`
	NISN      string    `json:"nisn" validate:"max=22"`
	FullName  string    `json:"full_name" validate:"required,min=1,max=70"`
	NickName  string    `json:"nick_name" validate:"omitempty,min=1,max=20"`
	Gender    string    `json:"gender" validate:"required,oneof=MALE FEMALE"`
	EntryYear string    `json:"entry_year" validate:"max=4"`
	Class     string    `json:"class" validate:"required,len=1,oneof=N 1 2 3 4 5 6 L"`
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
