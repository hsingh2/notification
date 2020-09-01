package cockroachdb

import (
	"context"
	"database/sql"
	"errors"

	template "cto-github.cisco.com/NFV-BU/xnotifservice/internal/notificationtemplate"
	"github.com/sirupsen/logrus"
)

//Error
var (
	ErrRepository = errors.New("unable to handle request")
)

type repository struct {
	db     *sql.DB
	logger *logrus.Logger
}

// New returns a concrete repository backed by CockroachDB
func New(db *sql.DB, logger *logrus.Logger) (template.Repository, error) {
	// return  repository
	return &repository{
		db:     db,
		logger: logger,
	}, nil
}

// CreateOrder inserts a new order and its order items into db
func (repo repository) CreateNotificationTemplate(ctx context.Context, req template.NotificationTemplate) (template.NotificationTemplate, error) {
	sql := `INSERT INTO notification_template (id, name, type, content, description)
	VALUES ($1,$2,$3,$4,$5) RETURNING id`

	result, err := repo.db.Exec(sql, req.ID, req.Name, req.Type, req.Content, req.Description)
	if err != nil {
		repo.logger.WithFields(logrus.Fields{"error": err}).Error("failed query create notificationtemplate")
		return template.NotificationTemplate{}, err
	}
	repo.logger.WithFields(logrus.Fields{"rows": result}).Info("create notificationtemplate")

	return req, nil
}

// CountNotificationTemplate ...
func (repo repository) CountNotificationTemplate(ctx context.Context) (int, error) {
	count := 0
	//return number of notification templates
	if err := repo.db.QueryRow("SELECT COUNT(*) from notification_template").Scan(&count); err != nil {
		repo.logger.WithFields(logrus.Fields{"error": err}).Error("failed query count notificationtemplate")
		return count, err
	}

	return count, nil
}

// GetNotificationTemplateByID ...
func (repo repository) GetNotificationTemplateByID(ctx context.Context, id string) (template.NotificationTemplate, error) {
	notificationTemplate := template.NotificationTemplate{}
	rows, err := repo.db.Query("SELECT * FROM notification_template where id = $1", id)
	if err != nil {
		repo.logger.WithFields(logrus.Fields{"error": err}).Error("failed query getbyID notificationtemplate")
		return notificationTemplate, err
	}
	//scan only one row
	for rows.Next() {
		if err := rows.Scan(&notificationTemplate); err != nil {
			repo.logger.WithFields(logrus.Fields{"error": err}).Error("failed scan row getbyID notificationtemplate")
			return notificationTemplate, err
		}
		break
	}
	return notificationTemplate, nil
}

//GetByPage returns all the Notification Template by page id
func (repo repository) GetByPage(context.Context) ([]template.NotificationTemplate, error) {
	return []template.NotificationTemplate{}, nil
}

//Update Notification Template
func (repo repository) UpdateNotificationTemplate(ctx context.Context, req template.NotificationTemplate) (template.NotificationTemplate, error) {
	sql := `UPDATE notification_template SET name=$2, type=$3, content=$4, description=$5 WHERE id=$1`

	if _, err := repo.db.Exec(sql, req.ID, req.Name, req.Type, req.Content, req.Description); err != nil {
		repo.logger.WithFields(logrus.Fields{"error": err}).Error("failed query update notificationtemplate")
		return template.NotificationTemplate{}, err
	}

	return req, nil
}

//Delete Notificaiton Template
func (repo repository) DeleteNotificationTemplate(ctx context.Context, id string) error {
	sql := `DELETE FROM notification_template WHERE id=$1`

	if _, err := repo.db.Exec(sql, id); err != nil {
		repo.logger.WithFields(logrus.Fields{"error": err}).Error("failed query delete notificationtemplate")
		return err
	}
	return nil
}

// Close implements DB.Close
func (repo repository) Close() error {
	return repo.db.Close()
}
