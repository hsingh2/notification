package admin

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

// Middleware describes a service middleware.
type Middleware func(service Service) Service

//LoggingMiddleware ...
func LoggingMiddleware(logger *logrus.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   Service
	logger *logrus.Logger
}

func (mw loggingMiddleware) Health(ctx context.Context) (status string, err error) {
	defer func(begin time.Time) {
		mw.logger.WithFields(logrus.Fields{"method": "Health",
			"request":  "/health",
			"err":      err,
			"response": status,
			"took":     time.Since(begin)}).Info("health middleware")
	}(time.Now())

	return mw.next.Health(ctx)
}

func (mw loggingMiddleware) Token(ctx context.Context) (token string, err error) {
	defer func(begin time.Time) {
		mw.logger.WithFields(logrus.Fields{"method": "Health",
			"request":  "/token",
			"err":      err,
			"response": token,
			"took":     time.Since(begin)}).Info("token middleware")
	}(time.Now())
	return mw.next.Token(ctx)
}
