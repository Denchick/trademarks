package store

import (
	"context"

	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/models"
)

// TrademarkRepository is a store for trademarks
type TrademarkRepository interface {
	FindTrademarkByName(context.Context, string) (*models.DBTrademark, error)
	FindSimilarTrademarks(ctx context.Context, name string) ([]*models.DBTrademark, error)
}