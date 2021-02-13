package controllers

import (
	"net/http"

	"github.com/denchick/trademarks/service"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// TrademarkController ...
type TrademarkController struct {
	services *service.Manager
}

// NewTrademark creates a new trademark controller
func NewTrademark(services *service.Manager) *TrademarkController {
	return &TrademarkController{services}
}

// Get returns trademark by ID
func (controller *TrademarkController) Get(c echo.Context) error {
	trademarks, err := controller.services.Trademark.GetTrademarks(
		c.QueryParam("name"), 
		c.QueryParam("similar") == "true",
	)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError, 
			errors.Wrap(err, "could not get trademark"),
		)
	}
	if len(trademarks) == 0 {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, trademarks)
}

