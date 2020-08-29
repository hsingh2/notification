package notificationtemplate

import (
	"context"
)

// NotificationTemplateService describes the Notification Template service.
type NotificationTemplateService interface {
	GetByID(context.Context, string) (NotificationTemplate, error)
	GetByPage(context.Context, string) ([]NotificationTemplate, error)
	Create(context.Context, NotificationTemplate) error
	Update(context.Context, NotificationTemplate) (NotificationTemplate, error)
	Delete(context.Context, string) error
	Count(context.Context) (int, error)
}
