package domain

import (
	"context"

	"github.com/openmesh/kitgen/models"
)

type StoreService interface {
	DeleteOrder(ctx context.Context, req DeleteOrderRequest) DeleteOrderResponse
	GetInventory(ctx context.Context, req GetInventoryRequest) GetInventoryResponse
	GetOrderByID(ctx context.Context, req GetOrderByIDRequest) GetOrderByIDResponse
	PlaceOrder(ctx context.Context, req PlaceOrderRequest) PlaceOrderResponse
}

type StoreService interface {
	DeleteOrder(ctx context.Context, orderId int64)
	GetInventory(ctx context.Context) map[string]int32
	GetOrderByID(ctx context.Context, orderId int64) models.Order
	PlaceOrder(ctx context.Context, body models.Order) models.Order
}
