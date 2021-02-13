package service

import "github.com/denchick/trademarks/models"

// TrademarkService is a store for trademarks
type TrademarkService interface {
	GetTrademarks(name string, similar bool) ([]*models.Trademark, error)
}