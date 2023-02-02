package requestlogger

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jclem/template-gofiber/requestcontext"
)

// Middleware is a Fiber middleware that logs the start and end of a request.
func Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		c.Locals("starttime", start)

		logStart(c)

		if err := c.Next(); err != nil {
			// We handle logging the end of the request in our error handler.
			// This is because we don't want to log until the error handler sets
			// the status code.
			return err
		}

		LogEnd(c)

		return nil
	}
}

func logStart(c *fiber.Ctx) {
	log := requestcontext.GetFromLocals(c).Logger
	start := c.Locals("starttime").(time.Time)

	log.Infow("request:start",
		"ts", start.Format(time.RFC3339),
		"method", c.Method(),
		"path", c.Path(),
		"requestID", c.Locals("requestid").(string),
		"ip", c.IP(),
	)
}

// LogEnd logs the end of a request. This function is public so that an error
// handler can log the end of a request in the event of an error.
func LogEnd(c *fiber.Ctx) {
	log := requestcontext.GetFromLocals(c).Logger
	start := c.Locals("starttime").(time.Time)

	log.Infow("request:end",
		"ts", time.Now().Format(time.RFC3339),
		"method", c.Method(),
		"route", c.Route().Path,
		"path", c.Path(),
		"requestID", c.Locals("requestid").(string),
		"ip", c.IP(),
		"status", c.Response().StatusCode(),
		"durationMs", time.Since(start).Milliseconds(),
	)
}
