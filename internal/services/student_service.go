package service

import (
	"github.com/dvvnFrtn/sisima/internal/config"
	model "github.com/dvvnFrtn/sisima/internal/models"
	"github.com/google/uuid"
)

type StudentService interface {
	Create(student *model.Student) error
	FindAllPaginated(page, limit int, sort, order string) ([]model.Student, int64, error)
	FindAll() ([]model.Student, error)
	FindDetailById(id uuid.UUID) (*model.Student, error)
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

func (s *studentService) FindAllPaginated(page, limit int, sort, order string) ([]model.Student, int64, error) {
	var students []model.Student
	var total int64
	config.DB.Model(&model.Student{}).Count(&total)
	err := config.DB.Limit(limit).Offset((page - 1) * limit).Order(sort + " " + order).Find(&students).Error
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
