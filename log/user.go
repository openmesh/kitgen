package log

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

func UserLoggingMiddleware(logger log.Logger) domain.UserServiceMiddleware {
	return func(next domain.UserService) domain.UserService {
		return userLoggingMiddleware{logger, next}
	}
}

type userLoggingMiddleware struct {
	logger log.Logger
	next   domain.UserService
}

func (mw *userLoggingMiddleware) CreateUser(ctx context.Context, body domain.User) (err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "create_user",
			"body", body,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return mw.next.CreateUser(ctx, body)
}

func (mw *userLoggingMiddleware) CreateUsersWithArrayInput(ctx context.Context, body []*domain.User) (err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "create_users_with_array_input",
			"body", body,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return mw.next.CreateUsersWithArrayInput(ctx, body)
}

func (mw *userLoggingMiddleware) CreateUsersWithListInput(ctx context.Context, body []*domain.User) (err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "create_users_with_list_input",
			"body", body,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return mw.next.CreateUsersWithListInput(ctx, body)
}

func (mw *userLoggingMiddleware) DeleteUser(ctx context.Context, username string) (err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "delete_user",
			"username", username,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return mw.next.DeleteUser(ctx, username)
}

func (mw *userLoggingMiddleware) GetUserByName(ctx context.Context, username string) (_ domain.User, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "get_user_by_name",
			"username", username,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return mw.next.GetUserByName(ctx, username)
}

func (mw *userLoggingMiddleware) LoginUser(ctx context.Context, password string, username string) (_ string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "login_user",
			"password", password,
			"username", username,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return mw.next.LoginUser(ctx, password, username)
}

func (mw *userLoggingMiddleware) LogoutUser(ctx context.Context) (err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "logout_user",
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return mw.next.LogoutUser(ctx)
}

func (mw *userLoggingMiddleware) UpdateUser(ctx context.Context, body domain.User, username string) (err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "update_user",
			"body", body,
			"username", username,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return mw.next.UpdateUser(ctx, body, username)
}
