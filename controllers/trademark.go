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

// Get returns file by ID
func (controller *TrademarkController) Get(c echo.Context) error {
	name := c.QueryParam("name")
	fuzzy := c.QueryParam("fuzzy")
	return c.String(http.StatusOK, "team:"+name+", member:"+fuzzy)
}
