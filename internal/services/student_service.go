package service

import (
	"errors"
	"time"

	"github.com/dvvnFrtn/sisima/internal/config"
	"github.com/dvvnFrtn/sisima/internal/enums"
	model "github.com/dvvnFrtn/sisima/internal/models"
	"github.com/google/uuid"
)

type StudentService interface {
	Create(student *model.Student) error
	FindAllPaginated(page, limit int, sort, order string, deleted bool) ([]model.Student, int64, error)
	FindAll() ([]model.Student, error)
	FindDetailById(id uuid.UUID) (*model.Student, error)
	DeleteById(id uuid.UUID, option enums.DeleteOptions) error
}

// struct
type studentService struct{}

// constructor
func NewStudentService() StudentService {
	return &studentService{}
}

// method
func (s *studentService) Create(student *model.Student) error {
	return config.DB.Create(student).Error
}

func (s *studentService) FindAll() ([]model.Student, error) {
	var students []model.Student
	err := config.DB.Find(&students).Error
	return students, err
}

func (s *studentService) FindAllPaginated(page, limit int, sort, order string, deleted bool) ([]model.Student, int64, error) {

	var students []model.Student
	var total int64
	var err error

	db := config.DB.Model(&model.Student{})

	if deleted {
		db = db.Where("deleted_at IS NOT NULL")
	} else {
		db = db.Where("deleted_at IS NULL")
	}

	if err = db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err = db.
		Limit(limit).
		Offset((page - 1) * limit).
		Order(sort + " " + order).
		Find(&students).Error

	if err != nil {
		return nil, 0, err
	}

	return students, total, nil
}

func (s *studentService) FindDetailById(id uuid.UUID) (*model.Student, error) {
	var student model.Student
	err := config.DB.First(&student, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (s *studentService) DeleteById(id uuid.UUID, option enums.DeleteOptions) error {
	switch option {

	case enums.Normal:
		tx := config.DB.Model(&model.Student{}).Where("id = ? AND deleted_at IS NULL", id).Update("deleted_at", time.Now())

		if tx.Error != nil {
			return tx.Error
		}
		if tx.RowsAffected == 0 {
			return errors.New("not-found-or-already-deleted")
		}
		return nil

	// not complete
	case enums.Rollback:
		tx := config.DB.Model(&model.Student{}).Where("id = ?", id).Update("deleted_at", nil)

		if tx.Error != nil {
			return tx.Error
		}
		if tx.RowsAffected == 0 {
			return errors.New("not-deleted-or-not-found")
		}
		return nil

	// not tested
	case enums.Hard:
		tx := config.DB.Unscoped().Where("id = ?", id).Delete(&model.Student{})

		if tx.Error != nil {
			return tx.Error
		}
		if tx.RowsAffected == 0 {
			return errors.New("not-found")
		}
		return nil

	default:
		return errors.New("invalid-option")
	}
}
