package admin

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints holds all Go kit endpoints for the Notification Template service.
type Endpoints struct {
	Health endpoint.Endpoint
	Token  endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints for the Notification Template service.
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Health: makeHealthEndpoint(s),
		Token:  makeTokenEndpoint(s),
	}
}

func makeHealthEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		message, err := s.Health(ctx)
		return HealthAdminTemplateResponse{Status: message, Err: err}, nil
	}
}

func makeTokenEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		message, err := s.Token(ctx)
		return TokenAdminTemplateResponse{Message: message, Err: err}, nil
	}
}
