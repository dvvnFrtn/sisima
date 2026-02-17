package handler

import (
	"fmt"
	"time"

	service "github.com/dvvnFrtn/sisima/internal/services"
	"github.com/gofiber/fiber/v3"
)

type APIInfo struct {
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Description string   `json:"description"`
	Developers  []string `json:"developers"`
	Timestamp   string   `json:"timestamp"`
	Copyright   string   `json:"copyright"`
}

func IndexHandler(c fiber.Ctx) error {
	info := APIInfo{
		Name:        "SISIMA",
		Version:     "1.0",
		Description: "Sistem Informasi Minhajussalam",
		Developers:  []string{"Achmed Hibatillah", "M. Rizki Fajar"},
		Timestamp:   time.Now().Format(time.RFC3339),
		Copyright:   "© 2026 SISIMA. All rights reserved.",
	}

	svc := service.NewBillingService()
	if err := svc.CreateBillingType(service.CreateBillingTypeRequest{
		Name:          "SPP",
		DefaultAmount: 50000,
		Recurring: &struct {
			Interval      string
			IntervalCount int64
		}{"MONTH", 1},
	}); err != nil {
		fmt.Println(err)
	}

	return c.JSON(info)
}
