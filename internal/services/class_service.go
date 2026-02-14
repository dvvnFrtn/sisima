package service

import "errors"

type ClassService interface {
	FindAllClassInYear(year uint8) error
}

// struct
type classService struct{}

// constructor
func NewClassService() ClassService {
	return &classService{}
}

// method
func (s *classService) FindAllClassInYear(year uint8) error {
	err := errors.New("error")
	return err
}
