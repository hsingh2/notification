package transport

import (
	"context"

	template "cto-github.cisco.com/NFV-BU/xnotifservice/internal/notificationtemplate"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints holds all Go kit endpoints for the Notification Template service.
type Endpoints struct {
	Create  endpoint.Endpoint
	GetByID endpoint.Endpoint
	Count   endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints for the Notification Template service.
func MakeEndpoints(s template.NotificationTemplateService) Endpoints {
	return Endpoints{
		Create:  makeCreateEndpoint(s),
		GetByID: makeGetByIDEndpoint(s),
		Count:   makeCountEndpoint(s),
	}
}

func makeCreateEndpoint(s template.NotificationTemplateService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateNotificationTemplateRequest) // type assertion
		err := s.Create(ctx, req.NotificationTemplate)
		return CreateNotificationTemplateResponse{Err: err}, nil
	}
}

func makeGetByIDEndpoint(s template.NotificationTemplateService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetNotificationTemplateByIDRequest)
		response, err := s.GetByID(ctx, req.ID)
		return GetNotificationTemplateByIDResponse{Template: response, Err: err}, nil
	}
}

func makeCountEndpoint(s template.NotificationTemplateService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		count, err := s.Count(ctx)
		return count, err
	}
}