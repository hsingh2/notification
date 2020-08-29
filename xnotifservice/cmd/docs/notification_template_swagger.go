package docs

import (
	"cto-github.cisco.com/NFV-BU/xnotifservice/internal/notificationtemplate/transport"
)

// swagger:route POST /api/v1/notificationTemplates notification-tag idOfFoobarEndpoint
// Notification Create Endpoint.
// responses:
//   200: notificationTemplates

// This text will appear as description of your response body.
// swagger:response foobarResponse
type foobarResponseWrapper struct {
	// in:body
	Body transport.CreateNotificationTemplateResponse
}

// swagger:parameters idOfFoobarEndpoint
type foobarParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body transport.CreateNotificationTemplateRequest
}
