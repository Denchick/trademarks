package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/store"
)

// TrademarkController ...
type TrademarkController struct {
	store *store.Store
}

// NewTrademark creates a new trademark controller
func NewTrademark(store *store.Store) *TrademarkController {
	return &TrademarkController{
		store: store,
	}
}

// Get returns trademark by ID
func (controller *TrademarkController) Get(c echo.Context) error {
	name := c.QueryParam("name")
	useFuzzy := c.QueryParam("searchFuzzy")
	trademark, err := controller.store.Trademark.FindTrademarkByName(c.Request().Context(), name, useFuzzy == "true")

	if err != nil || trademark == nil { // TODO impove error handling
		return echo.NewHTTPError(http.StatusNotFound, "Could not get trademark")
	}

	return c.JSON(http.StatusOK, trademark.ToTrademark())
}
