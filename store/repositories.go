package store

import (
	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/models"
)

// TrademarkRepository is a store for trademarks
type TrademarkRepository interface {
	FindTrademarkByName(name string) (*models.DBTrademark, error)
	FindSimilarTrademarks(name string) ([]*models.DBTrademark, error)
}
