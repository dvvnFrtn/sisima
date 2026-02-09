package logger

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func newRequestID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func HTTPLogger() fiber.Handler {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = os.Getenv("FIBER_ENV")
	}
	if env == "" {
		env = "development"
	}

	isDev := env == "development"

	return func(c *fiber.Ctx) error {
		start := time.Now()
		requestID := newRequestID()

		// expose ke handler lain & client
		c.Locals("request_id", requestID)
		c.Set("X-Request-ID", requestID)

		reqHeaders, _ := json.Marshal(c.GetReqHeaders())
		reqBody := string(c.Body())

		err := c.Next()

		latency := time.Since(start)
		status := c.Response().StatusCode()

		// =====================
		// TERMINAL LOG
		// =====================
		log.Printf(
			"%s | %3d | %8s | %15s | %-6s | %-32s | %s",
			time.Now().Format("15:04:05"),
			status,
			fmt.Sprintf("%.3fms", float64(latency.Microseconds())/1000),
			c.IP(),
			c.Method(),
			requestID,
			c.Path(),
		)

		// =====================
		// SQLITE (DEV ONLY)
		// =====================
		if isDev && db != nil {
			resHeaders := string(c.Response().Header.Header())
			resBody := string(c.Response().Body())

			_, dbErr := db.Exec(`
				INSERT INTO http_logs (
					request_id, ip, method, path, status, latency_us,
					req_headers, req_body, res_headers, res_body
				)
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
				requestID,
				c.IP(),
				c.Method(),
				c.Path(),
				status,
				latency.Microseconds(),
				string(reqHeaders),
				reqBody,
				resHeaders,
				resBody,
			)

			if dbErr != nil {
				log.Printf("[LOGGER][DB_ERROR] %v", dbErr)
			}
		}

		return err
	}
}
