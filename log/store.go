package log

import (
	"github.com/go-kit/kit/log"
)

func StoreLoggingMiddleware(logger log.Logger) domain.StoreServiceMiddleware {
	return func(next domain.StoreService) domain.StoreService {
		return storeLoggingMiddleware{logger, next}
	}
}
