package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/jclem/template-gofiber/appcontext"
	"github.com/jclem/template-gofiber/config"
	"github.com/jclem/template-gofiber/logger"
	"github.com/jclem/template-gofiber/meta"
	"github.com/jclem/template-gofiber/requestcontext"
	"github.com/jclem/template-gofiber/requestlogger"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	logger, err := logger.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	appctx, err := appcontext.NewWithOptions(
		appcontext.WithConfig(cfg),
		appcontext.WithLogger(logger),
	)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{ErrorHandler: errorHandler})
	app.Use(recover.New(recover.Config{EnableStackTrace: cfg.IsDev()}))
	app.Use(requestid.New())
	app.Use(requestcontext.Middleware(appctx))
	app.Use(requestlogger.Middleware())
	app.Get("/meta/healthcheck", meta.Healthcheck())

	log.Fatal(app.Listen(cfg.Addr()))
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	requestcontext.GetFromLocals(ctx).Logger.Errorf("error handling request: %s", err)
	handleErr := fiber.DefaultErrorHandler(ctx, err)
	requestlogger.LogEnd(ctx)

	return handleErr
}
