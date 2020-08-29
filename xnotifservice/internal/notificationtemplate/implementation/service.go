package implementation

import (
	"context"
	"database/sql"
	"errors"

	template "cto-github.cisco.com/NFV-BU/xnotifservice/internal/notificationtemplate"
	"github.com/sirupsen/logrus"
)

//Error
var (
	ErrNotificationNotFound = errors.New("order not found")
	ErrCmdRepository        = errors.New("unable to command repository")
	ErrQueryRepository      = errors.New("unable to query repository")
)

// service implements the Order Service
type service struct {
	repository template.Repository
	logger     *logrus.Logger
}

// NewTemplateService ...
func NewTemplateService(rep template.Repository, logger *logrus.Logger) template.NotificationTemplateService {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

// Create makes a notification template
func (s *service) Create(ctx context.Context, nTemplate template.NotificationTemplate) error {
	s.logger.Info("create service")
	if err := s.repository.CreateNotificationTemplate(ctx, nTemplate); err != nil {
		s.logger.WithFields(logrus.Fields{"error": err}).Error("create notification template repository return error")
		return ErrCmdRepository
	}

	return nil
}

// GetByID returns an notification template given by id
func (s *service) GetByID(ctx context.Context, id string) (template.NotificationTemplate, error) {
	s.logger.Info("getbyID service")

	template, err := s.repository.GetNotificationTemplateByID(ctx, id)
	if err != nil {
		s.logger.WithFields(logrus.Fields{"error": err}).Error("get notification template byid return error")

		if err == sql.ErrNoRows {
			return template, ErrNotificationNotFound
		}
		return template, ErrQueryRepository
	}
	return template, nil
}

// Count returns notification template counts
func (s *service) Count(ctx context.Context) (int, error) {
	count, err := s.repository.CountNotificationTemplate(ctx)
	if err != nil {
		s.logger.Error(err)
		return count, ErrCmdRepository
	}
	return count, nil
}

//GetByPage return notification template by page
func (s *service) GetByPage(context.Context, string) ([]template.NotificationTemplate, error) {
	return []template.NotificationTemplate{}, nil
}

//Update updates the notification template
func (s *service) Update(context.Context, template.NotificationTemplate) (template.NotificationTemplate, error) {
	return template.NotificationTemplate{}, nil
}

//Delete deletes notification template
func (s *service) Delete(context.Context, string) error {
	return nil
}
