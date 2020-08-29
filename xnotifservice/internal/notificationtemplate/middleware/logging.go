package middleware

import (
	"context"

	"time"

	template "cto-github.cisco.com/NFV-BU/xnotifservice/internal/notificationtemplate"
	log "github.com/sirupsen/logrus"
)

//LoggingMiddleware ...
func LoggingMiddleware(logger *log.Logger) Middleware {
	return func(next template.NotificationTemplateService) template.NotificationTemplateService {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   template.NotificationTemplateService
	logger *log.Logger
}

func (mw loggingMiddleware) Create(ctx context.Context, template template.NotificationTemplate) (err error) {
	defer func(begin time.Time) {
		mw.logger.WithFields(log.Fields{"method": "Create",
			"request": template,
			"err":     err,
			"took":    time.Since(begin)}).Info("create notificationtemplate middleware")
	}(time.Now())

	return mw.next.Create(ctx, template)
}

func (mw loggingMiddleware) GetByID(ctx context.Context, id string) (template template.NotificationTemplate, err error) {
	defer func(begin time.Time) {
		mw.logger.WithFields(log.Fields{"method": "GetByID",
			"request":  id,
			"err":      err,
			"response": template,
			"took":     time.Since(begin)}).Info("getbyid notificationtemplate middleware")
	}(time.Now())
	return mw.next.GetByID(ctx, id)
}

func (mw loggingMiddleware) Count(ctx context.Context) (count int, err error) {
	defer func(begin time.Time) {
		mw.logger.WithFields(log.Fields{"method": "Count",
			"request":  "count all",
			"err":      err,
			"response": count,
			"took":     time.Since(begin)}).Info("count notificationtemplate middleware")
	}(time.Now())

	return mw.next.Count(ctx)
}

func (mw loggingMiddleware) Delete(ctx context.Context, id string) (err error) {
	defer func(begin time.Time) {
		mw.logger.WithFields(log.Fields{"method": "Delete",
			"request": id,
			"err":     err,
			"took":    time.Since(begin)}).Info("delete notificationtemplate middleware")
	}(time.Now())
	return mw.next.Delete(ctx, id)
}

func (mw loggingMiddleware) GetByPage(ctx context.Context, id string) (templates []template.NotificationTemplate, err error) {
	defer func(begin time.Time) {
		mw.logger.WithFields(log.Fields{"method": "GetByPage",
			"request":  id,
			"err":      err,
			"response": templates,
			"took":     time.Since(begin)}).Info("getbypage notificationtemplate middleware")
	}(time.Now())
	return mw.next.GetByPage(ctx, id)
}

func (mw loggingMiddleware) Update(ctx context.Context, template template.NotificationTemplate) (updated template.NotificationTemplate, err error) {
	defer func(begin time.Time) {
		mw.logger.WithFields(log.Fields{"method": "Update",
			"request":  template,
			"err":      err,
			"response": updated,
			"took":     time.Since(begin)}).Info("update notificationtemplate middleware")
	}(time.Now())
	return mw.next.Update(ctx, template)
}
