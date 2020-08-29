package cockroachdb

import (
	"context"

	template "cto-github.cisco.com/NFV-BU/xnotifservice/internal/notificationtemplate"
)

//MockRepository ...
type mockClient struct {
}

//MockRepository ...
func MockRepository() template.Repository {
	return &mockClient{}
}

func (mck mockClient) GetNotificationTemplateByID(context.Context, string) (template.NotificationTemplate, error) {
	return template.NotificationTemplate{Name: "yasemin"}, nil
}
func (mck mockClient) CreateNotificationTemplate(context.Context, template.NotificationTemplate) error {
	return nil
}
func (mck mockClient) CountNotificationTemplate(context.Context) (int, error) {
	return 100, nil
}

func (mck mockClient) GetByPage(context.Context) ([]template.NotificationTemplate, error) {
	return []template.NotificationTemplate{}, nil
}
func (mck mockClient) UpdateNotificationTemplate(context.Context, template.NotificationTemplate) (template.NotificationTemplate, error) {
	return template.NotificationTemplate{}, nil
}

func (mck mockClient) DeleteNotificationTemplate(context.Context, string) error {
	return nil
}
