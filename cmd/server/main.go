package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/controllers"
	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/store"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Init repository store (with mysql/postgresql inside)
	store, err := store.New()
	if err != nil {
		return errors.Wrap(err, "store.New failed")
	}

	// Init controllers
	trademarksController := controllers.NewTrademark(store)

	e := echo.New()

	// API version 1
	v1 := e.Group("/v1")

	// Trademark routes
	trademarkRoutes := v1.Group("/trademarks")
	trademarkRoutes.GET("", trademarksController.Get)

	// ping routes
	pingRoutes := v1.Group("/ping")
	pingRoutes.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))

	return nil
}
