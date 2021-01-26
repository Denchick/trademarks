package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/logger"
	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/models"
	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/store"
)

// TrademarkController ...
type TrademarkController struct {
	store  *store.Store
	logger *logger.Logger
}

// NewTrademark creates a new trademark controller
func NewTrademark(store *store.Store, logger *logger.Logger) *TrademarkController {
	return &TrademarkController{
		store:  store,
		logger: logger,
	}
}

// Get returns trademark by ID
func (controller *TrademarkController) Get(c echo.Context) error { // TODO create OpenAPI specification
	// TODO search should be case insensitive
	name := c.QueryParam("name")
	fuzzily := c.QueryParam("fuzzily")
	trademarks, err := controller.getTrademarks(c, name, fuzzily == "true")

	if err != nil {
		controller.logger.Err(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, controller.convertTrademarks(trademarks))
}

func (controller *TrademarkController) getTrademarks(c echo.Context, name string, fuzzily bool) ([]*models.DBTrademark, error) {
	if fuzzily {
		return controller.store.Trademark.FindSimilarTrademarks(c.Request().Context(), name)
	}
	trademark, err := controller.store.Trademark.FindTrademarkByName(c.Request().Context(), name)
	return []*models.DBTrademark{trademark}, err
}

func (controller *TrademarkController) convertTrademarks(oldTrademarks []*models.DBTrademark) []*models.Trademark {
	var newTrademarks []*models.Trademark
	for _, trademark := range oldTrademarks {
		newTrademarks = append(newTrademarks, trademark.ToTrademark())
	}
	return newTrademarks
}
