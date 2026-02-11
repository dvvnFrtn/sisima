package handler

import (
	"time"

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

	return c.JSON(info)
}
