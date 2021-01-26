package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/config"
	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/controllers"
	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/logger"
	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/store"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	store, err := store.New()
	if err != nil {
		return errors.Wrap(err, "store.New failed")
	}

	logger := logger.Get(config.Get().LogLevel)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group("/v1")

	trademarksController := controllers.NewTrademark(store, logger)
	trademarkRoutes := v1.Group("/trademarks")
	trademarkRoutes.GET("", trademarksController.Get)

	pingRoutes := v1.Group("/_ping")
	pingRoutes.GET("", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.Logger.Print(e.Start(":1323")) // TODO read port from the config

	return nil
}
