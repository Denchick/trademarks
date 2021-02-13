package main

import (
	"log"
	"net/http"
	"time"

	"github.com/denchick/trademarks/config"
	"github.com/denchick/trademarks/controllers"
	"github.com/denchick/trademarks/service"
	"github.com/denchick/trademarks/store"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg := config.Get()

	// Init store
	store, err := store.NewStore()
	if err != nil {
		return errors.Wrap(err, "store.NewStore")
	}
	
	// Init manager
	manager, err := service.NewManager(store)
	if err != nil {
		return errors.Wrap(err, "manager.NewManager")
	}

	// Initialize Echo instance
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Init controllers
	v1 := e.Group("/v1")

	trademarksController := controllers.NewTrademarkController(manager)
	trademarkRoutes := v1.Group("/trademarks")
	trademarkRoutes.GET("", trademarksController.Get)

	pingRoutes := v1.Group("/_ping")
	pingRoutes.GET("", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	// Start server
	s := &http.Server{
		Addr:         cfg.HTTPAddr,
		ReadTimeout:  30 * time.Minute,
		WriteTimeout: 30 * time.Minute,
	}
	e.Logger.Print(e.StartServer(s))

	return nil
}
