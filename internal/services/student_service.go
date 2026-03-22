package service

import (
	"strings"

	"github.com/dvvnFrtn/sisima/internal/config"
	model "github.com/dvvnFrtn/sisima/internal/models"
	"github.com/google/uuid"
)

type StudentService interface {
	Create(student *model.Student) error
	FindSomeLimited(page, limit int, sort, order string, filterGender model.Gender, filterClass string, keyword string) ([]model.StudentWithTotalBills, int64, error)
	FindAll() ([]model.Student, error)
	FindDetailById(id uuid.UUID) (*model.Student, error)
	GetIdsByName(full_name string) ([]uuid.UUID, error)
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

func (s *studentService) FindSomeLimited(
	page, limit int,
	sort, order string,
	filterGender model.Gender,
	filterClass string,
	keyword string,
) ([]model.StudentWithTotalBills, int64, error) {
	var students []model.StudentWithTotalBills
	var total int64

	query := config.DB.Model(&model.Student{}).
		Select("students.id, students.full_name, students.nis, students.nisn, students.gender, students.class, count(billings.amount) as total_tagihan").
		Joins("LEFT JOIN billings ON billings.student_id = students.id").
		Group("students.id, students.full_name, students.nis, students.nisn, students.gender, students.class")

	if filterGender != "" {
		query = query.Where("students.gender = ?", filterGender)
	}

	if filterClass != "" {
		query = query.Where("students.class = ?", filterClass)
	}

	if keyword != "" {
		kw := strings.ToLower(keyword)
		query = query.Where(
			"LOWER(students.full_name) LIKE ? OR students.nis LIKE ? OR students.nisn LIKE ?",
			"%"+kw+"%", "%"+kw+"%", "%"+kw+"%",
		)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.
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

func (s *studentService) GetIdsByName(full_name string) ([]uuid.UUID, error) {
	var ids []uuid.UUID

	err := config.DB.
		Model(&model.Student{}).
		Where("full_name = ?", full_name).
		Pluck("id", &ids).Error
	if err != nil {
		return nil, err
	}

	return ids, nil
}
