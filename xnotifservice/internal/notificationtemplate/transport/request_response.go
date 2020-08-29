package transport

import (
	template "cto-github.cisco.com/NFV-BU/xnotifservice/internal/notificationtemplate"
)

//CreateNotificationTemplateRequest ...
type CreateNotificationTemplateRequest struct {
	NotificationTemplate template.NotificationTemplate
}

//CreateNotificationTemplateResponse ...
type CreateNotificationTemplateResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error,omitempty"`
}

//GetNotificationTemplateByIDRequest ...
type GetNotificationTemplateByIDRequest struct {
	ID string
}

//GetNotificationTemplateByIDResponse ...
type GetNotificationTemplateByIDResponse struct {
	Template template.NotificationTemplate `json:"notificationTemplate"`
	Err      error                         `json:"error,omitempty"`
}

//CountNotificationTemplateResponse ...
type CountNotificationTemplateResponse struct {
	Count int   `json:"count,omitempty"`
	Err   error `json:"error,omitempty"`
}
