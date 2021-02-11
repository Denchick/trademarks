package store

import (
	"github.com/denchick/trademarks/models"
)

// TrademarkRepository is a store for trademarks
type TrademarkRepository interface {
	FindTrademarkByName(name string) (*models.DBTrademark, error)
	FindSimilarTrademarks(name string) ([]*models.DBTrademark, error)
}
