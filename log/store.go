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

func (mw *storeLoggingMiddleware) DeleteOrder(ctx context.Context, orderId int64) (err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "delete_order",
			"orderId", orderId,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return mw.next.DeleteOrder(ctx, orderId)
}

func (mw *storeLoggingMiddleware) GetInventory(ctx context.Context) (_ map[string]int32, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "get_inventory",
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return mw.next.GetInventory(ctx)
}

func (mw *storeLoggingMiddleware) GetOrderByID(ctx context.Context, orderId int64) (_ domain.Order, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "get_order_by_id",
			"orderId", orderId,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return mw.next.GetOrderByID(ctx, orderId)
}

func (mw *storeLoggingMiddleware) PlaceOrder(ctx context.Context, body domain.Order) (_ domain.Order, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "place_order",
			"body", body,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return mw.next.PlaceOrder(ctx, body)
}
