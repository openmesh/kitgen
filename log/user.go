package log

import (
	"github.com/go-kit/kit/log"
)

func UserLoggingMiddleware(logger log.Logger) domain.UserServiceMiddleware {
	return func(next domain.UserService) domain.UserService {
		return userLoggingMiddleware{logger, next}
	}
}
