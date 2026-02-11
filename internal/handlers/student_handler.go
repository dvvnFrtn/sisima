package handler

import (
	"strconv"

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
	const (
		defaultPage  = 1
		defaultLimit = 10
		defaultSort  = "full_name"
		defaultOrder = "ASC"
	)

	pageStr := c.Query("page", strconv.Itoa(defaultPage))
	limitStr := c.Query("limit", strconv.Itoa(defaultLimit))
	sort := c.Query("sort", defaultSort)
	order := c.Query("order", defaultOrder)

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = defaultPage
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = defaultLimit
	}

	students, total, err := h.service.FindAllPaginated(page, limit, sort, order)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"title":  "INTERNAL_ERROR",
			"errors": err,
		})
	}

	response := dto.NewPagination(make([]interface{}, len(students)), page, limit, total)

	for i, v := range students {
		var student model.Student
		student = v
		response.Data[i] = dto.ToStudentResponse(&student)
	}

	return c.Status(200).JSON(response)
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
