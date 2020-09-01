package docs

import (
	"cto-github.cisco.com/NFV-BU/xnotifservice/internal/notificationtemplate"
	"cto-github.cisco.com/NFV-BU/xnotifservice/internal/notificationtemplate/transport"
)

// swagger:route GET /api/v1/notificationTemplates/{id} NotificationTemplate getNotificationTemplate
// API to get one NotificationTemplate by id
// responses:
//   200: getResponse

// swagger:route POST /api/v1/notificationTemplates NotificationTemplate createNotificationTemplate
// API to create a NotificationTemplate
// responses:
//   200: createResponse

// Success response for the create Notification Template.
// swagger:response createResponse
type createResponseWrapper struct {
	// in:body
	Body transport.CreateNotificationTemplateResponse
}

// swagger:parameters createNotificationTemplate
type createParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body transport.CreateNotificationTemplateRequest
}

// swagger:response getResponse
type getResponseWrapper struct {
	// in:body
	Body notificationtemplate.NotificationTemplate
}

// A NotificationTemplateID parameter model.
//
// This is used for operations that want the ID of an NotificationTemplateID in the path
// swagger:parameters getNotificationTemplate
type NotificationTemplateID struct {
	// The ID of the pet
	//
	// in: path
	// required: true
	ID string `json:"id"`
}
