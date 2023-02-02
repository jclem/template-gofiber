package requestcontext

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jclem/template-gofiber/appcontext"
)

var localKeyName = "jclem/template-gofiber:requestcontext"

// Middleware is a Fiber middleware that attaches a copy of the app context to the request.
func Middleware(ctx *appcontext.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if _, err := AttachContext(c, ctx); err != nil {
			return err
		}

		return c.Next()
	}
}

// AttachContext attaches a copy of the app context to the request and returns the copy.
func AttachContext(c *fiber.Ctx, ctx *appcontext.Context) (*appcontext.Context, error) {
	logger := ctx.Logger.With("requestid", c.Locals("requestid"))

	ctx, err := ctx.WithOptions(
		appcontext.WithContext(c.Context()),
		appcontext.WithLogger(logger),
	)
	if err != nil {
		return nil, err
	}

	c.Locals(localKeyName, ctx)
	c.Context().SetUserValue(localKeyName, ctx)

	return ctx, nil
}

type hasLocals interface {
	Locals(key any, value ...any) any
}

// GetFromLocals returns the app context from a Fiber context.
func GetFromLocals(c hasLocals) *appcontext.Context {
	return c.Locals(localKeyName).(*appcontext.Context)
}
