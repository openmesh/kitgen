package domain

import (
	"context"
)

type StoreService interface {
	DeleteOrder(ctx context.Context, orderID int64) error
	GetInventory(ctx context.Context) (map[string]int32, error)
	GetOrderByID(ctx context.Context, orderID int64) (Order, error)
	PlaceOrder(ctx context.Context, body Order) (Order, error)
}
