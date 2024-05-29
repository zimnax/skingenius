package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	applog "skingenius/logger"

	"github.com/google/uuid"
)

func FiberMiddleware(a *fiber.App) {
	a.Use(
		cors.New(),
		logger.New(),
		injectTransactionId(),
	)
}

func injectTransactionId() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals(applog.TransactionId, uuid.New().String())
		return c.Next()
	}
}
