package implementation

import (
	"context"
	"database/sql"
	"errors"

	template "cto-github.cisco.com/NFV-BU/xnotifservice/internal/notificationtemplate"
	"github.com/sirupsen/logrus"

	"github.com/google/uuid"
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
func (s *service) Create(ctx context.Context, nTemplate template.NotificationTemplate) (template.NotificationTemplate, error) {
	s.logger.Info("create service")
	//add id
	nTemplate.ID = uuid.New().String()

	response, err := s.repository.CreateNotificationTemplate(ctx, nTemplate)
	if err != nil {
		s.logger.WithFields(logrus.Fields{"error": err}).Error("create notification template repository return error")
		return template.NotificationTemplate{}, ErrCmdRepository
	}

	return response, nil
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
func (s *service) Update(ctx context.Context, request template.NotificationTemplate) (template.NotificationTemplate, error) {
	response, err := s.repository.UpdateNotificationTemplate(ctx, request)
	if err != nil {
		s.logger.WithFields(logrus.Fields{"error": err}).Error("update notification template return error")
		return template.NotificationTemplate{}, ErrCmdRepository
	}
	return response, nil
}

//Delete deletes notification template
func (s *service) Delete(ctx context.Context, id string) error {
	if err := s.repository.DeleteNotificationTemplate(ctx, id); err != nil {
		s.logger.WithFields(logrus.Fields{"error": err}).Error("delete notification template return error")
		return ErrCmdRepository
	}

	return nil
}
