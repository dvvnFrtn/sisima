package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;column:id"`
	NIS       string    `gorm:"type:varchar(22);unique;not null;column:nis"`
	NISN      string    `gorm:"type:varchar(22);unique;not null;column:nisn"`
	FullName  string    `gorm:"type:varchar(70);not null;column:full_name"`
	NickName  string    `gorm:"type:varchar(20);column:nick_name"`
	Gender    Gender    `gorm:"type:gender_enum;not null;column:gender"`
	EntryYear string    `gorm:"type:char(4);not null;column:entry_year"`
	Class     string    `gorm:"type:char(1);not null;column:class"`

	CreatedAt time.Time      `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (s *Student) BeforeCreate(tx *gorm.DB) error {
	s.ID = uuid.New()
	return nil
}
