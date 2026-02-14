package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Class struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;column:id"`
	Number         int       `gorm:"type:integer;not null;column:number"`
	AcademicYearID uuid.UUID `gorm:"type:uuid;not null;column:academic_year_id"`

	AcademicYear AcademicYear `gorm:"foreignKey:AcademicYearID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;`
}

func (s *Class) BeforeCreateClassId(tx *gorm.DB) error {
	s.ID = uuid.New()
	return nil
}
