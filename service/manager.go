package service

import (
	"github.com/denchick/trademarks/service/services"
	"github.com/denchick/trademarks/store"
	"github.com/pkg/errors"
)

// Manager is just a collection of all services we have in the project
type Manager struct {
	Trademark TrademarkService
}

// NewManager creates new service manager
func NewManager(store *store.Store) (*Manager, error) {
	if store == nil {
		return nil, errors.New("No store provided")
	}
	return &Manager{
		Trademark: services.NewTrademarkService(store),
	}, nil
}
