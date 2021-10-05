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

func (mw *createUserLoggingMiddleware) CreateUser(ctx context.Context, body User) (err error) {
	defer func(beging time.Time) {
		_ = mw.logger.Log(
			"method", "create_user",
			"body", body,
		)
	}()
}

func (mw *createUsersWithArrayInputLoggingMiddleware) CreateUsersWithArrayInput(ctx context.Context, body User) (err error) {
	defer func(beging time.Time) {
		_ = mw.logger.Log(
			"method", "create_users_with_array_input",
			"body", body,
		)
	}()
}

func (mw *createUsersWithListInputLoggingMiddleware) CreateUsersWithListInput(ctx context.Context, body User) (err error) {
	defer func(beging time.Time) {
		_ = mw.logger.Log(
			"method", "create_users_with_list_input",
			"body", body,
		)
	}()
}

func (mw *deleteUserLoggingMiddleware) DeleteUser(ctx context.Context, username string) (err error) {
	defer func(beging time.Time) {
		_ = mw.logger.Log(
			"method", "delete_user",
			"username", username,
		)
	}()
}

func (mw *getUserByNameLoggingMiddleware) GetUserByName(ctx context.Context, username string) (getUserByNameOK domain.User, err error) {
	defer func(beging time.Time) {
		_ = mw.logger.Log(
			"method", "get_user_by_name",
			"username", username,
		)
	}()
}

func (mw *loginUserLoggingMiddleware) LoginUser(ctx context.Context, password string, username string) (loginUserOK string, err error) {
	defer func(beging time.Time) {
		_ = mw.logger.Log(
			"method", "login_user",
			"password", password,
			"username", username,
		)
	}()
}

func (mw *logoutUserLoggingMiddleware) LogoutUser(ctx context.Context) (err error) {
	defer func(beging time.Time) {
		_ = mw.logger.Log(
			"method", "logout_user",
		)
	}()
}

func (mw *updateUserLoggingMiddleware) UpdateUser(ctx context.Context, body User, username string) (err error) {
	defer func(beging time.Time) {
		_ = mw.logger.Log(
			"method", "update_user",
			"body", body,
			"username", username,
		)
	}()
}
