package admin

import (
	"context"

	"github.com/sirupsen/logrus"
)

// service implements the Admin Service
type service struct {
	logger *logrus.Logger
}

// Service ...
type Service interface {
	Health(context.Context) (string, error)
	Token(context.Context) (string, error)
}

// NewAdminService ...
func NewAdminService(logger *logrus.Logger) Service {
	return &service{logger}
}

func (svc service) Health(context.Context) (string, error) {
	return "UP", nil
}

func (svc service) Token(context.Context) (string, error) {
	return "token", nil
}
