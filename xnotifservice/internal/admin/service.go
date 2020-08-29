package admin

import (
	"context"

	"github.com/go-kit/kit/log"
)

// service implements the Admin Service
type service struct {
	logger *log.Logger
}

// Service ...
type Service interface {
	Health(context.Context) (string, error)
	Token(context.Context) (string, error)
}

// NewAdminService ...
func NewAdminService(logger *log.Logger) Service {
	return &service{logger}
}

func (svc service) Health(context.Context) (string, error) {
	return "UP", nil
}

func (svc service) Token(context.Context) (string, error) {
	return "token", nil
}
