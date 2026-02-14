package model

import (
	"github.com/google/uuid"
)

type AcademicYear struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey;column:id"`
	Start   int16     `gorm:"type:smallint;unique;column:start"`
	Name    string    `gorm:"type:varchar(70);column:name"`
	Classes []Class   `gorm:"foreignKey:AcademicYearID"`
}
