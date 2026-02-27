package handler

import (
	"github.com/dvvnFrtn/sisima/internal/dto"
	service "github.com/dvvnFrtn/sisima/internal/services"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type billingHandler struct {
	service service.BillingService
}

func NewBillingHandler(s service.BillingService) *billingHandler {
	return &billingHandler{service: s}
}

func (bh *billingHandler) CreateBillingType(c fiber.Ctx) error {
	var req dto.CreateBillingTypeRequest
	if err := c.Bind().Body(&req); err != nil {
		// return c.Status(fiber.StatusBadRequest).
		// 	JSON(fiber.Map{
		// 		"title": "INVALID_REQUEST",
		// 	})
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.NewExceptionResponse(dto.InvalidRequest, nil))
	}

	if err := dto.Validate(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			JSON(dto.NewExceptionResponse(dto.ValidationErr, err.Errors))
	}

	if err := bh.service.CreateBillingType(req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.NewExceptionResponse(dto.InternalErr, err.Error()))
	}

	return c.Status(fiber.StatusCreated).End()
}

func (bh *billingHandler) UpdateBillingType(c fiber.Ctx) error {
	var req dto.UpdateBillingTypeRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.NewExceptionResponse(dto.InvalidRequest, nil))

	}

	if err := dto.Validate(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			JSON(dto.NewExceptionResponse(dto.ValidationErr, err.Errors))

	}

	btID := c.Params("billing_type_id")
	if err := bh.service.UpdateBillingType(uuid.MustParse(btID), req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.NewExceptionResponse(dto.InternalErr, err.Error()))

	}

	return c.Status(fiber.StatusOK).End()
}

func (bh *billingHandler) GetAllBillingType(c fiber.Ctx) error {
	resp, err := bh.service.GetAllBillingType()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.NewExceptionResponse(dto.InternalErr, err.Error()))

	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (bh *billingHandler) GetBillingType(c fiber.Ctx) error {
	btID := c.Params("billing_type_id")
	resp, err := bh.service.GetBillingType(uuid.MustParse(btID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.NewExceptionResponse(dto.InternalErr, err.Error()))

	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (bh *billingHandler) CreateBilling(c fiber.Ctx) error {
	var req dto.CreateBillingRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.NewExceptionResponse(dto.InvalidRequest, nil))

	}

	if err := dto.Validate(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			JSON(dto.NewExceptionResponse(dto.ValidationErr, err.Errors))

	}

	if err := bh.service.CreateBilling(req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.NewExceptionResponse(dto.InternalErr, err.Error()))

	}

	return c.Status(fiber.StatusCreated).End()
}

func (bh *billingHandler) GetAllBilling(c fiber.Ctx) error {
	resp, err := bh.service.GetAllBilling()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.NewExceptionResponse(dto.InternalErr, err.Error()))

	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (bh *billingHandler) GetBilling(c fiber.Ctx) error {
	bID := c.Params("billing_id")
	resp, err := bh.service.GetBilling(uuid.MustParse(bID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.NewExceptionResponse(dto.InternalErr, err.Error()))

	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
