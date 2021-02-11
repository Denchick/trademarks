package store

import (
	"github.com/denchick/trademarks/models"
)

// TrademarkRepository is a store for trademarks
type TrademarkRepository interface {
	FindByName(name string) ([]*models.DBTrademark, error)
	FindSimilar(name string) ([]*models.DBTrademark, error)
}
