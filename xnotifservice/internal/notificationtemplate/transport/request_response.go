package transport

import (
	template "cto-github.cisco.com/NFV-BU/xnotifservice/internal/notificationtemplate"
)

//CreateNotificationTemplateRequest ...
type CreateNotificationTemplateRequest struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Content     string `json:"content"`
	Description string `json:"description"`
}

//CreateNotificationTemplateResponse ...
type CreateNotificationTemplateResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Content     string `json:"content"`
	Description string `json:"description"`
	Err         error  `json:"error,omitempty"`
}

//GetNotificationTemplateByIDRequest ...
type GetNotificationTemplateByIDRequest struct {
	ID string
}

//GetNotificationTemplateByIDResponse ...
type GetNotificationTemplateByIDResponse struct {
	template.NotificationTemplate
	Err error `json:"error,omitempty"`
}

//CountNotificationTemplateResponse ...
type CountNotificationTemplateResponse struct {
	Count int   `json:"count,omitempty"`
	Err   error `json:"error,omitempty"`
}

//UpdateNotificationTemplateRequest ...
type UpdateNotificationTemplateRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Content     string `json:"content"`
	Description string `json:"description"`
}
