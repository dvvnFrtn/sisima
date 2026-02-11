package handler

import (
	"github.com/dvvnFrtn/sisima/internal/dto"
	model "github.com/dvvnFrtn/sisima/internal/models"
	service "github.com/dvvnFrtn/sisima/internal/services"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

// struct
type studentHandler struct {
	service service.StudentService
}

// constructor
func NewStudentHandler(s service.StudentService) *studentHandler {
	return &studentHandler{service: s}
}

// method
func (h *studentHandler) FindAllPaginated(c fiber.Ctx) error {
	data, err := h.service.FindAllPaginated(0, 3)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"title":  "INTERNAL_ERROR",
			"errors": err,
		})
	}

	return c.Status(200).JSON(data)
}

func (h *studentHandler) FindDetailById(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"title": "INVALID_ID",
		})
	}

	student, err := h.service.FindDetailById(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"title": "NOT_FOUND",
		})
	}
	response := dto.ToStudentResponse(student)
	return c.Status(200).JSON(response)
}

func (h *studentHandler) Create(c fiber.Ctx) error {
	var req dto.CreateStudentRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"title": "INVALID_REQUEST",
		})
	}

	if err := dto.Validate(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"title":  "VALIDATION_ERROR",
			"errors": err.Errors,
		})
	}

	student := &model.Student{
		NIS:       req.NIS,
		NISN:      req.NISN,
		FullName:  req.FullName,
		NickName:  req.NickName,
		Gender:    model.Gender(req.Gender),
		EntryYear: req.EntryYear,
		Class:     req.Class,
	}

	if err := h.service.Create(student); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"title":  "INTERNAL_ERROR",
			"errors": err.Error(),
		})
	}

	response := dto.ToStudentResponse(student)

	return c.Status(200).JSON(response)
}
