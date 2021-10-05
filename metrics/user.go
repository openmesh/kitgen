package metrics

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"
)

func UserMetricsMiddleware(
	requestCount metrics.Counter,
	errorCount metrics.Counter,
	requestDuration metrics.Histogram,
) domain.UserServiceMiddleware {
	return func(next domain.UserService) domain.UserService {
		return userMetricsMiddleware{requestCount, errorCount, requestDuration, next}
	}
}

type userMetricsMiddleware struct {
	requestCount    metrics.Counter
	errorCount      metrics.Counter
	requestDuration metrics.Histogram
	next            domain.UserService
}

func (mw *userMetricsMiddleware) CreateUser(ctx context.Context, body domain.User) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "create_user"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	return mw.next.CreateUser(ctx, body)
}

func (mw *userMetricsMiddleware) CreateUsersWithArrayInput(ctx context.Context, body []*domain.User) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "create_users_with_array_input"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	return mw.next.CreateUsersWithArrayInput(ctx, body)
}

func (mw *userMetricsMiddleware) CreateUsersWithListInput(ctx context.Context, body []*domain.User) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "create_users_with_list_input"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	return mw.next.CreateUsersWithListInput(ctx, body)
}

func (mw *userMetricsMiddleware) DeleteUser(ctx context.Context, username string) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "delete_user"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	return mw.next.DeleteUser(ctx, username)
}

func (mw *userMetricsMiddleware) GetUserByName(ctx context.Context, username string) (_ domain.User, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get_user_by_name"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	return mw.next.GetUserByName(ctx, username)
}

func (mw *userMetricsMiddleware) LoginUser(ctx context.Context, password string, username string) (_ string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "login_user"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	return mw.next.LoginUser(ctx, password, username)
}

func (mw *userMetricsMiddleware) LogoutUser(ctx context.Context) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "logout_user"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	return mw.next.LogoutUser(ctx)
}

func (mw *userMetricsMiddleware) UpdateUser(ctx context.Context, body domain.User, username string) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "update_user"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	return mw.next.UpdateUser(ctx, body, username)
}
