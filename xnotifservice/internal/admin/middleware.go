package admin

import (
	"context"

	"github.com/go-kit/kit/log"

	"time"
)

// Middleware describes a service middleware.
type Middleware func(service Service) Service

//LoggingMiddleware ...
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   Service
	logger log.Logger
}

func (mw loggingMiddleware) Health(ctx context.Context) (status string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Health", "took", "healthstatus", status, "err", err, time.Since(begin))
	}(time.Now())
	return mw.next.Health(ctx)
}

func (mw loggingMiddleware) Token(ctx context.Context) (token string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Health", "took", "token", token, "err", err, time.Since(begin))
	}(time.Now())
	return mw.next.Token(ctx)
}
