package server

import (
	"fmt"
	"log"

	"github.com/alfianazizi/personal-finance/config"
	"github.com/alfianazizi/personal-finance/database"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type fiberServer struct {
	app  *fiber.App
	db   database.Database
	conf *config.Config
}

func NewFiberServer(conf *config.Config, db database.Database) Server {
	fiberApp := fiber.New()

	return &fiberServer{
		app:  fiberApp,
		db:   db,
		conf: conf,
	}
}

func (f *fiberServer) Start() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	f.app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))

	f.app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	serverUrl := fmt.Sprintf(":%d", f.conf.Server.Port)
	log.Fatal(f.app.Listen(serverUrl))

}
