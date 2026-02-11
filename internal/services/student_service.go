package service

import (
	"github.com/dvvnFrtn/sisima/internal/config"
	"github.com/dvvnFrtn/sisima/internal/dto"
	model "github.com/dvvnFrtn/sisima/internal/models"
	"github.com/google/uuid"
)

type StudentService interface {
	Create(student *model.Student) error
	FindAllPaginated(page int, limit int) (*dto.Pagination, error)
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

func (s *studentService) FindAllPaginated(page int, limit int) (*dto.Pagination, error) {
	var students []model.Student
	err := config.DB.Limit(limit).Offset((page - 1) * limit).Find(&students).Error
	if err != nil {
		return nil, err
	}

	data := dto.Pagination{
		Data: make([]interface{}, len(students)),
		Meta: dto.Meta{
			Page: page,
		},
	}

	for i, v := range students {
		data.Data[i] = v
	}

	return &data, nil
}

func (s *studentService) FindDetailById(id uuid.UUID) (*model.Student, error) {
	var student model.Student
	err := config.DB.First(&student, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}
