package notificationtemplate

import (
	"context"
	"errors"
)

//Errors
var (
	ErrNotificationNotFound = errors.New("order not found")
	ErrCmdRepository        = errors.New("unable to command repository")
	ErrQueryRepository      = errors.New("unable to query repository")
)

//NotificationTemplate ...
type NotificationTemplate struct {
	Name string `json:"name"`
}

//Repository ...
type Repository interface {
	GetNotificationTemplateByID(context.Context, string) (NotificationTemplate, error)
	GetByPage(context.Context) ([]NotificationTemplate, error)
	CreateNotificationTemplate(context.Context, NotificationTemplate) error
	UpdateNotificationTemplate(context.Context, NotificationTemplate) (NotificationTemplate, error)
	CountNotificationTemplate(context.Context) (int, error)
	DeleteNotificationTemplate(context.Context, string) error
}
