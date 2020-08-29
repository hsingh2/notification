package middleware

import (
	template "cto-github.cisco.com/NFV-BU/xnotifservice/internal/notificationtemplate"
)

// Middleware describes a service middleware.
type Middleware func(service template.NotificationTemplateService) template.NotificationTemplateService
