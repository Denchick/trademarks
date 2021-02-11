package controllers

import (
	"net/http"

	"github.com/denchick/trademarks/models"
	"github.com/denchick/trademarks/store"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// TrademarkController ...
type TrademarkController struct {
	store  *store.Store
}

// NewTrademark creates a new trademark controller
func NewTrademark(store *store.Store) *TrademarkController {
	return &TrademarkController{
		store:  store,
	}
}

// Get returns trademark by ID
func (ctr *TrademarkController) Get(c echo.Context) error {
	name := c.QueryParam("name")
	fuzzily := c.QueryParam("fuzzily")
	trademarks, err := ctr.getTrademarks(c, name, fuzzily == "true")

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not get trademark"))
	}
	if len(trademarks) == 0 {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, ctr.convert(trademarks))
}

func (ctr *TrademarkController) getTrademarks(c echo.Context, name string, fuzzily bool) ([]*models.DBTrademark, error) {
	if fuzzily {
		return ctr.store.Trademark.FindSimilar(name)
	}
	return ctr.store.Trademark.FindByName(name)
}

func (ctr *TrademarkController) convert(oldTrademarks []*models.DBTrademark) []*models.Trademark {
	var newTrademarks []*models.Trademark
	for _, trademark := range oldTrademarks {
		newTrademarks = append(newTrademarks, trademark.ToTrademark())
	}
	return newTrademarks
}
