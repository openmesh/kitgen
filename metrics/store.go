package metrics

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"
)

func StoreMetricsMiddleware(
	requestCount metrics.Counter,
	errorCount metrics.Counter,
	requestDuration metrics.Histogram,
) domain.StoreServiceMiddleware {
	return func(next domain.StoreService) domain.StoreService {
		return storeMetricsMiddleware{requestCount, errorCount, requestDuration, next}
	}
}

type storeMetricsMiddleware struct {
	requestCount    metrics.Counter
	errorCount      metrics.Counter
	requestDuration metrics.Histogram
	next            domain.StoreService
}

func (mw *storeMetricsMiddleware) DeleteOrder(ctx context.Context, orderId int64) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "delete_order"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	return mw.next.DeleteOrder(ctx, orderId)
}

func (mw *storeMetricsMiddleware) GetInventory(ctx context.Context) (_ map[string]int32, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get_inventory"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	return mw.next.GetInventory(ctx)
}

func (mw *storeMetricsMiddleware) GetOrderByID(ctx context.Context, orderId int64) (_ domain.Order, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get_order_by_id"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	return mw.next.GetOrderByID(ctx, orderId)
}

func (mw *storeMetricsMiddleware) PlaceOrder(ctx context.Context, body domain.Order) (_ domain.Order, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "place_order"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	return mw.next.PlaceOrder(ctx, body)
}
