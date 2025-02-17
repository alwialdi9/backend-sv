package middlewares

import (
	"time"

	"github.com/alwialdi9/backend-sv/internal/utils"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func LoggerMiddleware(c *fiber.Ctx) error {
	// Log request
	start := time.Now()

	// Process request
	err := c.Next()

	// Log request details
	utils.LogInfo("Incoming Request", log.Fields{
		"method":  c.Method(),
		"path":    c.Path(),
		"status":  c.Response().StatusCode(),
		"latency": time.Since(start),
	})

	return err
}
