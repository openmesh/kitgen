package log

import (
	"github.com/go-kit/kit/log"
)

func PetLoggingMiddleware(logger log.Logger) domain.PetServiceMiddleware {
	return func(next domain.PetService) domain.PetService {
		return petLoggingMiddleware{logger, next}
	}
}
