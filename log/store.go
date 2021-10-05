package log

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

func StoreLoggingMiddleware(logger log.Logger) domain.StoreServiceMiddleware {
	return func(next domain.StoreService) domain.StoreService {
		return storeLoggingMiddleware{logger, next}
	}
}

type storeLoggingMiddleware struct {
	logger log.Logger
	next   domain.StoreService
}

func (mw *deleteOrderLoggingMiddleware) DeleteOrder(ctx context.Context, orderID int64) (err error) {
	defer func(beging time.Time) {
		_ = mw.logger.Log(
			"method", "delete_order",
			"orderId", orderId,
		)
	}()
}

func (mw *getInventoryLoggingMiddleware) GetInventory(ctx context.Context) (getInventoryOK map[string]int32, err error) {
	defer func(beging time.Time) {
		_ = mw.logger.Log(
			"method", "get_inventory",
		)
	}()
}

func (mw *getOrderByIdLoggingMiddleware) GetOrderByID(ctx context.Context, orderID int64) (getOrderByIdOK domain.Order, err error) {
	defer func(beging time.Time) {
		_ = mw.logger.Log(
			"method", "get_order_by_id",
			"orderId", orderId,
		)
	}()
}

func (mw *placeOrderLoggingMiddleware) PlaceOrder(ctx context.Context, body Order) (placeOrderOK domain.Order, err error) {
	defer func(beging time.Time) {
		_ = mw.logger.Log(
			"method", "place_order",
			"body", body,
		)
	}()
}
