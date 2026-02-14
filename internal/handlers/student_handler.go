package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dvvnFrtn/sisima/internal/dto"
	"github.com/dvvnFrtn/sisima/internal/enums"
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
	pageStr := c.Query("page", "1")
	limitStr := c.Query("limit", "10")
	sort := c.Query("sort", "full_name")
	order := c.Query("order", "ASC")
	deletedStr := c.Query("deleted", "false")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(limitStr)
	if limit < 1 {
		limit = 10
	}

	deleted := false
	if strings.ToLower(deletedStr) == "true" {
		deleted = true
	}

	allowedSort := map[string]bool{"full_name": true}
	allowedOrder := map[string]bool{"ASC": true, "DESC": true}
	if !allowedSort[sort] {
		sort = "full_name"
	}
	order = strings.ToUpper(order)
	if !allowedOrder[order] {
		order = "ASC"
	}

	students, total, err := h.service.FindAllPaginated(page, limit, sort, order, deleted)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"title":  "INTERNAL_ERROR",
			"errors": err.Error(),
		})
	}

	response := dto.NewPagination(make([]interface{}, len(students)), page, limit, total)
	for i, v := range students {
		response.Data[i] = dto.ToStudentResponse(&v)
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

func (h *studentHandler) Delete(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"title": "INVALID_ID",
		})
	}

	optionQuery := c.Query("option")
	option, err := enums.ParseDeleteOption(optionQuery)
	if err != nil {
		option = enums.Normal
	}

	err = h.service.DeleteById(id, option)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"title":  "INTERNAL_ERROR",
			"errors": err,
		})
	}

	return c.SendStatus(204)
}
