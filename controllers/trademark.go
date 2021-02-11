package controllers

import (
	"net/http"

	"github.com/denchick/trademarks/logger"
	"github.com/denchick/trademarks/models"
	"github.com/denchick/trademarks/store"
	"github.com/labstack/echo/v4"
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
func (ctr *TrademarkController) Get(c echo.Context) error {
	name := c.QueryParam("name")
	fuzzily := c.QueryParam("fuzzily")
	trademarks, err := ctr.getTrademarks(c, name, fuzzily == "true")

	if err != nil {
		ctr.logger.Err(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, ctr.convertTrademarks(trademarks))
}

func (ctr *TrademarkController) getTrademarks(c echo.Context, name string, fuzzily bool) ([]*models.DBTrademark, error) {
	if fuzzily {
		return ctr.store.Trademark.FindSimilarTrademarks(name)
	}
	trademark, err := ctr.store.Trademark.FindTrademarkByName(name)
	return []*models.DBTrademark{trademark}, err
}

func (ctr *TrademarkController) convertTrademarks(oldTrademarks []*models.DBTrademark) []*models.Trademark {
	var newTrademarks []*models.Trademark
	for _, trademark := range oldTrademarks {
		newTrademarks = append(newTrademarks, trademark.ToTrademark())
	}
	return newTrademarks
}
